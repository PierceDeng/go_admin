package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dsn := "dpz:Dpz18229744479!@tcp(192.168.30.128:3306)/plain?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	DB = db

}

func CloseDB() {

	sqlDB, _ := DB.DB()
	if sqlDB != nil {
		err := sqlDB.Close()
		if err != nil {
			return
		}
	}
}
