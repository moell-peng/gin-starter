package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moell/config"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	conf := config.Get()
	db, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err == nil {
		if conf.SQLDebug {
			DB = db.Debug()
		} else {
			DB = db
		}

		return db, err
	}
	return nil, err
}
