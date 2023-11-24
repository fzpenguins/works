package config

import (
	"Memorandum/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func InitDB(dsn string) {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("err = ", err)
		return
	}
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Task{})
}
