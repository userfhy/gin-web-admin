package model

import (
    "fmt"
    "gin-test/utils/setting"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "log"
    "time"
)

var db *gorm.DB

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

    // 设置连接池中的最大闲置连接数。
    db.DB().SetMaxIdleConns(10)
    
    // 设置数据库的最大连接数量。
    db.DB().SetMaxOpenConns(100)
    
    // 设置连接的最大可复用时间。
    db.DB().SetConnMaxLifetime(time.Hour)

    db.SingularTable(true)

    // 不存在 创建表
    //if ! db.HasTable(&Report{}) {
    //   log.Println("不存在上报表，开始创建！")
    //   db.CreateTable(&Report{})
    //}

    // 自动迁移表
    db.AutoMigrate(&Report{})

    gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
        return setting.DatabaseSetting.TablePrefix + defaultTableName
    }
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