package model

import (
	"database/sql/driver"
	"fmt"
	"gin-web-admin/utils/setting"
	"log"
	"os"
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
	res := db.Delete(tableStruct)
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

func GetTotal(tableStruct interface{}, whereSql string, values []interface{}) (int64, error) {
	var count int64
	if err := db.Model(tableStruct).Where(whereSql, values...).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

//func GetTotal(maps interface{}) (int, error) {
//    var count int
//    if err := db.Model(&Auth{}).Where(maps).Count(&count).Error; err != nil {
//        return 0, err
//    }
//
//    return count, nil
//}
//
//// GetTestUsers gets a list of users based on paging constraints
//func GetList(pageNum int, pageSize int, maps interface{}) ([]*interface{}, error) {
//    var user [] *Auth
//    err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&user).Error
//    if err != nil && err != gorm.ErrRecordNotFound {
//        return nil, err
//    }
//
//    return user, nil
//}

func BuildCondition(where map[string]interface{}) (whereSql string, values []interface{}, err error) {
	for key, value := range where {
		conditionKey := strings.Split(key, " ")
		if len(conditionKey) > 2 {
			return "", nil, fmt.Errorf("" +
				"map构建的条件格式不对，类似于'age >'")
		}
		if whereSql != "" {
			whereSql += " AND "
		}
		switch len(conditionKey) {
		case 1:
			whereSql += fmt.Sprint(conditionKey[0], " = ?")
			values = append(values, value)
			break
		case 2:
			field := conditionKey[0]
			switch conditionKey[1] {
			case "=":
				whereSql += fmt.Sprint(field, " = ?")
				values = append(values, value)
				break
			case ">":
				whereSql += fmt.Sprint(field, " > ?")
				values = append(values, value)
				break
			case ">=":
				whereSql += fmt.Sprint(field, " >= ?")
				values = append(values, value)
				break
			case "<":
				whereSql += fmt.Sprint(field, " < ?")
				values = append(values, value)
				break
			case "<=":
				whereSql += fmt.Sprint(field, " <= ?")
				values = append(values, value)
				break
			case "in":
				whereSql += fmt.Sprint(field, " in (?)")
				values = append(values, value)
				break
			case "like":
				whereSql += fmt.Sprint(field, " like ?")
				values = append(values, value)
				break
			case "<>":
				whereSql += fmt.Sprint(field, " != ?")
				values = append(values, value)
				break
			case "!=":
				whereSql += fmt.Sprint(field, " != ?")
				values = append(values, value)
				break
			}
			break
		}
	}
	return
}
