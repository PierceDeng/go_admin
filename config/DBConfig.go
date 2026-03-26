package config

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dsn := viper.GetString("database.dsn")
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
