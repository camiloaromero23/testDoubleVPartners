package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "root"
const DB_NAME = "tickets_db"

var DB *gorm.DB

func InitDatabase(DB_HOST string, DB_PORT string) *gorm.DB {
	DB = ConnectDataBase(DB_HOST, DB_PORT)
	return DB
}

func ConnectDataBase(DB_HOST string, DB_PORT string) *gorm.DB {
	//var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	//database.AutoMigrate(&models.Ticket{})
	return database

	//sql.DB{}
}
