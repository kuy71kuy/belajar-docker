package config

import (
	"fmt"
	c "project/constants"
	"project/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() {
	initDB()
	initialMigration()
}

func initDB() {
	config := model.Config{
		DbUsername: c.LoadEnv("DB_USERNAME"),
		DbPassword: c.LoadEnv("DB_PASSWORD"),
		DbPort:     c.LoadEnv("DB_PORT"),
		DbHost:     c.LoadEnv("DB_HOST"),
		DbName:     c.LoadEnv("DB_NAME"),
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DbUsername,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	initialMigration()
}

func initialMigration() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Book{})
}
