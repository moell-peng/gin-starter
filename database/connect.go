package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gin-gorm-example/models"
	"gin-gorm-example/config"
)


var  DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	conf := config.Get()
	db, err := gorm.Open(mysql.Open(conf.DSN),&gorm.Config{})

	if err == nil {
		DB = db
		err := db.AutoMigrate(&models.AdminUser{})
		return db, err
	}
	return nil, err
}