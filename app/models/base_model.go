package model

import (
	"database/sql/driver"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"gin-web-admin/utils/logging"
	"gin-web-admin/utils/setting"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db          *gorm.DB
	TablePrefix string
)

// JSONTime format json time field by myself
type JSONTime struct {
	time.Time
}

type PaginateStruct struct {
	PageNum  int
	PageSize int
}

type BaseModel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt JSONTime  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt JSONTime  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *JSONTime `sql:"index" json:"deleted_at"`
}

type BaseModelNoId struct {
	CreatedAt JSONTime  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt JSONTime  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *JSONTime `sql:"index" json:"deleted_at"`
}

func Setup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=20s",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name,
	)

	var newLogger logger.Interface

	log.Println("Connecting to database...")
	if gin.Mode() == gin.DebugMode {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,       // True - Don't include params in the SQL log
				Colorful:                  true,
			},
		)
	}

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})

	if err != nil {
		log.Fatalf("Base models.Setup err: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("DB err: %v", err)
	}
	// 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	/*    gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	      return TablePrefix + defaultTableName
	  }*/

	TablePrefix = setting.DatabaseSetting.TablePrefix

	// 不存在 创建表
	//if ! db.HasTable(&Report{}) {
	//   log.Println("不存在上报表，开始创建！")
	//   db.CreateTable(&Report{})
	//}

	// 自动迁移表
	if err := db.AutoMigrate(
		&Report{},
		&Auth{},
		&JwtBlacklist{},
		&Role{},
		// &CasbinRule{},
		&Menu{},
		&Dept{},
		&RoleMenu{},
		&SiteContent{},
		&SiteCategory{},
		&SiteContentCategory{},
		&SiteTag{},
		&SiteContentTag{},
		&DictType{},
		&DictData{},
		&AuditLog{},
	); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}
}

func DBClose() {
	sql, _ := db.DB()
	if err := sql.Close(); err != nil {
		logging.Warnf("db close failed: %v", err)
	}
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`null`), nil
	}
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t JSONTime) IsZero() bool {
	return t.Time.IsZero()
}

func (t JSONTime) Format(layout string) string {
	return t.Time.Format(layout)
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v any) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

/*func (v BaseModel) BeforeCreate(scope *gorm.Scope) error {
   scope.SetColumn("created_at", time.Now())
   scope.SetColumn("updated_at", time.Now())
   return nil
}

func (v BaseModel) BeforeUpdate(scope *gorm.Scope) error {
   scope.SetColumn("updated_at", time.Now())
   return nil
}*/

func SoftDelete(tableStruct any) (int64, error) {
	res := db.Model(tableStruct).Update("deleted_at", time.Now())
	if err := res.Error; err != nil {
		return 0, err
	}
	return res.RowsAffected, nil
}

// 新增字段名验证函数
func validateUpdateFields(updates map[string]any) error {
	for field := range updates {
		if err := validateColumnName(field); err != nil {
			return fmt.Errorf("invalid field name %q: %v", field, err)
		}
	}
	return nil
}

func Update(tableStruct any, where map[string]any, updates map[string]any) (int64, error) {
	if len(updates) == 0 {
		return 0, fmt.Errorf("updates cannot be empty")
	}

	// 验证更新字段名
	if err := validateUpdateFields(updates); err != nil {
		return 0, err
	}

	// 安全构建WHERE条件
	dbData, err := BuildCondition(db.Model(tableStruct), where)
	if err != nil {
		return 0, err
	}

	res := dbData.Updates(updates)
	if err := res.Error; err != nil {
		return 0, err
	}
	return res.RowsAffected, nil
}

func GetTotal(tableStruct any, where map[string]any) (int64, error) {
	var count int64
	dbData, err := BuildCondition(db.Model(tableStruct), where)
	if err != nil {
		return 0, err
	}

	if err := dbData.Model(tableStruct).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// SafeOperators contains all allowed SQL operators
var SafeOperators = map[string]string{
	"=":    "=",
	">":    ">",
	">=":   ">=",
	"<":    "<",
	"<=":   "<=",
	"!=":   "!=",
	"<>":   "<>",
	"in":   "IN",
	"like": "LIKE",
	"is":   "IS",
}

// BuildCondition builds SQL conditions safely
func BuildCondition(d *gorm.DB, where map[string]any) (*gorm.DB, error) {
	for field, value := range where {
		parts := strings.Fields(field)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid condition format: %q (expected 'field operator')", field)
		}

		column, operator := parts[0], strings.ToLower(parts[1])

		// 字段名验证
		if err := validateColumnName(column); err != nil {
			return nil, fmt.Errorf("invalid column name %q: %v", column, err)
		}

		// 操作符白名单验证
		safeOperator, ok := SafeOperators[operator]
		if !ok {
			return nil, fmt.Errorf("unsupported operator: %q", operator)
		}

		quotedColumn := quoteColumnName(column)

		// 安全处理不同操作符
		switch safeOperator {
		case "IN":
			if err := validateInClauseValues(value); err != nil {
				return nil, fmt.Errorf("invalid IN clause values: %v", err)
			}
			d = d.Where(quotedColumn+" IN ?", value)

		case "LIKE":
			pattern, ok := value.(string)
			if !ok {
				return nil, fmt.Errorf("LIKE operator requires string value")
			}
			escapedPattern := escapeLikePattern(pattern)
			d = d.Where(quotedColumn+" LIKE ? ESCAPE '\\'", escapedPattern)

		case "IS":
			// 严格处理IS操作符，只允许NULL/NOT NULL
			if value == nil {
				d = d.Where(quotedColumn + " IS NULL")
			} else if s, ok := value.(string); ok {
				switch strings.ToUpper(s) {
				case "NULL":
					d = d.Where(quotedColumn + " IS NULL")
				case "NOT NULL":
					d = d.Where(quotedColumn + " IS NOT NULL")
				default:
					return nil, fmt.Errorf("invalid value for IS operator: %q", s)
				}
			} else {
				return nil, fmt.Errorf("invalid type for IS operator value")
			}

		default:
			// 标准比较操作符
			d = d.Where(quotedColumn+" "+safeOperator+" ?", value)
		}
	}
	return d, nil
}

func quoteColumnName(name string) string {
	return "`" + name + "`"
}

func isLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

// 增强的字段名验证
func validateColumnName(name string) error {
	if len(name) == 0 || len(name) > 64 {
		return fmt.Errorf("column name must be 1-64 characters")
	}

	// 使用正则表达式验证会更严格，这里简化处理
	for i, c := range name {
		if i == 0 && !isLetter(c) {
			return fmt.Errorf("first character must be a letter")
		}
		if !isLetter(c) && !isDigit(c) && c != '_' {
			return fmt.Errorf("invalid character %q in column name", c)
		}
	}
	return nil
}

// validateInClauseValues validates values for IN clause
func validateInClauseValues(value any) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return fmt.Errorf("IN clause requires slice/array type")
	}

	if v.Len() == 0 {
		return fmt.Errorf("IN clause requires non-empty values")
	}

	if v.Len() > 1000 {
		return fmt.Errorf("IN clause exceeds maximum allowed values (1000)")
	}

	// 类型一致性检查（可选）
	// elemType := v.Type().Elem()
	// for i := 0; i < v.Len(); i++ {
	//     if v.Index(i).Type() != elemType {
	//         return fmt.Errorf("inconsistent element types in slice")
	//     }
	// }

	return nil
}

// escapeLikePattern escapes special characters in LIKE patterns
func escapeLikePattern(pattern string) string {
	escapeChar := "\\"
	replacer := strings.NewReplacer(
		escapeChar, escapeChar+escapeChar,
		"%", escapeChar+"%",
		"_", escapeChar+"_",
	)
	return replacer.Replace(pattern)
}
