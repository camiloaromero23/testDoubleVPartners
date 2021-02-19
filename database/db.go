package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME string
	DB_PORT string
	DB_HOST string
}

var DB *gorm.DB

func InitDatabase(dbConfig DatabaseConfig) *gorm.DB {
	DB = ConnectDataBase(dbConfig)
	return DB
}

func ConnectDataBase(dbConfig DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.DB_USERNAME,
		dbConfig.DB_PASSWORD,
		dbConfig.DB_HOST,
		dbConfig.DB_PORT,
		dbConfig.DB_NAME)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	return database
}
