package model

import (
    "database/sql/driver"
    "fmt"
    "gin-test/utils/setting"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "log"
    "time"
)

var (
    db *gorm.DB
    TablePrefix string
)

// JSONTime format json time field by myself
type JSONTime struct {
    time.Time
}

type BaseModel struct {
    ID        uint `gorm:"primary_key" json:"id"`
    CreatedAt JSONTime `gorm:"column:created_at" json:"created_at"`
    UpdatedAt JSONTime `gorm:"column:updated_at" json:"updated_at"`
    DeletedAt *JSONTime `sql:"index" json:"deleted_at"`
}

func Setup() {
    var err error
    db, err = gorm.Open(
        setting.DatabaseSetting.Type,
        fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=20s",
            setting.DatabaseSetting.User,
            setting.DatabaseSetting.Password,
            setting.DatabaseSetting.Host,
            setting.DatabaseSetting.Name,
        ),
    )

    if err != nil {
        log.Fatalf("Base models.Setup err: %v", err)
    }

    if setting.DatabaseSetting.EchoSql {
        db.LogMode(true)
    }

    // 设置连接池中的最大闲置连接数。
    db.DB().SetMaxIdleConns(10)
    
    // 设置数据库的最大连接数量。
    db.DB().SetMaxOpenConns(100)
    
    // 设置连接的最大可复用时间。
    db.DB().SetConnMaxLifetime(time.Hour)

    db.SingularTable(true)

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
    )
}

func TestDB() {
    var err error
    err = db.DB().Ping()
    if err != nil {
        log.Fatalf("DB ping err: %v", err)
    }
    
    // Scan
    type Result struct {
        Id int
        Name string
    }

    rows, err := db.Raw("SELECT id,name FROM t_user").Rows()
    defer rows.Close()
    
    var result Result
    for rows.Next() {
        //rows.Scan(&result)
        db.ScanRows(rows, &result)
        log.Println(result)
    }
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
    defer db.Close()
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

func Update(tableStruct interface{}, wheres map[string]interface{}, updates map[string]interface{}) (error, int64){
    res := db.Model(tableStruct).Where(wheres).Update(updates)
    if err := res.Error; err != nil {
        return err, 0
    }
    return nil, res.RowsAffected
}

func GetTotal(tableStruct interface{}, maps interface{}) (int, error) {
    var count int
    if err := db.Model(tableStruct).Where(maps).Count(&count).Error; err != nil {
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
