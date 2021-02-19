package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"testDoubleVPartners/controllers"
	"testDoubleVPartners/database"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	PORT := GetEnvVariable("PORT")
	DB_USERNAME := GetEnvVariable("DB_USERNAME")
	DB_PASSWORD := GetEnvVariable("DB_PASSWORD")
	DB_NAME := GetEnvVariable("DB_NAME")
	DB_PORT := GetEnvVariable("DB_PORT")
	DB_HOST := GetEnvVariable("DB_HOST")

	dbConfig := database.DatabaseConfig {
		DB_USERNAME,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
		DB_HOST,
	}

	fmt.Println("Listening on port: " + PORT)

	router := routerSetup(dbConfig)

	if err := router.Run(":" + PORT); err != nil {
		panic(err.Error())
	}
}

func GetEnvVariable(key string) string {
	if err := godotenv.Load("environment.env"); err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(key)
}

func routerSetup(dbConfig database.DatabaseConfig) *gin.Engine {
	router := gin.Default()

	ticketRepo := controllers.New(dbConfig)

	ticketRoutes := router.Group("/tickets")
	{
		ticketRoutes.POST("/", ticketRepo.CreateTicket)
		ticketRoutes.GET("/", ticketRepo.GetTickets)
		ticketRoutes.GET("/:id", ticketRepo.GetTicket)
		ticketRoutes.PUT("/:id", ticketRepo.UpdateTicket)
		ticketRoutes.DELETE("/:id", ticketRepo.DeleteTicket)
	}
	return router
}
