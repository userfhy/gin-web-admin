package model

import (
	"database/sql/driver"
	"fmt"
	"gin-web-admin/utils/setting"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

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

	if setting.DatabaseSetting.EchoSql {
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
	db.AutoMigrate(
		&Report{},
		&Auth{},
		&JwtBlacklist{},
		&Role{},
		// &CasbinRule{},
		&Menu{},
	)
}

func DBClose() {
	sql, _ := db.DB()
	sql.Close()
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`null`), nil
	}
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
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

func SoftDelete(tableStruct interface{}) (error, int64) {
	log.Println(tableStruct)
	res := db.Model(tableStruct).Update("deleted_at", time.Now())
	if err := res.Error; err != nil {
		return err, 0
	}
	return nil, res.RowsAffected
}

func Update(tableStruct interface{}, wheres map[string]interface{}, updates map[string]interface{}) (error, int64) {
	res := db.Model(tableStruct).Where(wheres).Updates(updates)
	if err := res.Error; err != nil {
		return err, 0
	}
	return nil, res.RowsAffected
}

func GetTotal(tableStruct interface{}, where map[string]interface{}) (int64, error) {
	var count int64
	var dbData, err = BuildCondition(db, where)

	if err != nil {
		fmt.Println("Error:", err)
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
func BuildCondition(d *gorm.DB, where map[string]interface{}) (*gorm.DB, error) {
	for field, value := range where {
		parts := strings.Fields(field)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid condition format: %q (expected 'field operator')", field)
		}

		column, operator := parts[0], strings.ToLower(parts[1])

		// Validate column name
		if err := validateColumnName(column); err != nil {
			return nil, fmt.Errorf("invalid column name %q: %v", column, err)
		}

		// Validate operator
		safeOperator, ok := SafeOperators[operator]
		if !ok {
			return nil, fmt.Errorf("unsupported operator: %q", operator)
		}

		// Handle different operators safely
		switch safeOperator {
		case "IN":
			if err := validateInClauseValues(value); err != nil {
				return nil, fmt.Errorf("invalid IN clause values: %v", err)
			}
			d = d.Where(column+" IN (?)", value)

		case "LIKE":
			pattern, ok := value.(string)
			if !ok {
				return nil, fmt.Errorf("LIKE operator requires string value")
			}
			escapedPattern := escapeLikePattern(pattern)
			d = d.Where(column+" LIKE ?", escapedPattern)

		case "IS":
			if value == nil {
				d = d.Where(column + " IS NULL")
			} else {
				d = d.Where(column+" = ?", value)
			}

		default:
			// Handle standard comparison operators
			d = d.Where(column+" "+safeOperator+" ?", value)
		}
	}

	return d, nil
}

// validateColumnName checks if the column name is safe
func validateColumnName(name string) error {
	if len(name) == 0 || len(name) > 64 { // MySQL's maximum identifier length
		return fmt.Errorf("column name length must be between 1 and 64 characters")
	}

	// Only allow alphanumeric characters and underscores
	for i, char := range name {
		if i == 0 {
			if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_') {
				return fmt.Errorf("column name must start with a letter or underscore")
			}
		} else {
			if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') ||
				(char >= '0' && char <= '9') || char == '_') {
				return fmt.Errorf("column name can only contain letters, numbers, and underscores")
			}
		}
	}
	return nil
}

// validateInClauseValues validates values for IN clause
func validateInClauseValues(value interface{}) error {
	// Check if the value is a slice
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return fmt.Errorf("IN clause requires slice or array")
	}

	// Check if the slice is not empty
	if v.Len() == 0 {
		return fmt.Errorf("IN clause requires non-empty slice")
	}

	// Maximum number of items in IN clause (adjust as needed)
	if v.Len() > 1000 {
		return fmt.Errorf("too many values in IN clause (max 1000)")
	}

	return nil
}

// escapeLikePattern escapes special characters in LIKE patterns
func escapeLikePattern(pattern string) string {
	pattern = strings.ReplaceAll(pattern, "%", "\\%")
	pattern = strings.ReplaceAll(pattern, "_", "\\_")
	return pattern
}
