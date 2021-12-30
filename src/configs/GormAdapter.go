package configs

import (
	"fmt"
	"github.com/WangYiwei-oss/jdframe/src/configinjector"
	"github.com/WangYiwei-oss/jdframe/src/configparser"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

func init() {
	configinjector.RegisterBean("MYSQL", NewGormAdapter)
}

var doOnceGorm sync.Once

type GormAdapter struct {
	*gorm.DB
}

var db *gorm.DB

func NewGormAdapter() *GormAdapter {
	var err error
	doOnceGorm.Do(func() {
		config := configparser.GlobalSettings["SQL_CONFIG"].(map[string]interface{})
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config["USER"].(string),
			config["PASSWORD"].(string),
			config["IP"].(string),
			config["PORT"].(string),
			config["DATABASE"].(string))
		//dsn := "root:123456@tcp(127.0.0.1:3306)/jdnotes?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalln("连接数据库失败")
		}
		mysqldb, err := db.DB()
		if err != nil {
			log.Fatalln("创建mysqldb失败")
		}
		mysqldb.SetMaxOpenConns(10)
		mysqldb.SetMaxIdleConns(5)
		log.Println("连接数据库成功")
	})
	return &GormAdapter{DB: db}
}
