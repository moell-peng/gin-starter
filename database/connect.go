package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"gin-gorm-example/models"
	"gin-gorm-example/config"
)


var  DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	conf := config.Get()
	db, err := gorm.Open("mysql", conf.DSN)

	if err == nil {
		db.DB().SetMaxIdleConns(conf.MaxIdleConn)
		DB = db
		db.AutoMigrate(&models.AdminUser{})
		return db, err
	}
	return nil, err
}