package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"testDoubleVPartners/controllers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	PORT := os.Getenv("PORT")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	fmt.Println(PORT + DB_HOST + DB_PORT)
	router := routerSetup(DB_HOST, DB_PORT)

	if err := router.Run(":" + PORT); err != nil {
		panic(err.Error())
	}
}

func routerSetup(DB_HOST string, DB_PORT string) *gin.Engine {
	router := gin.Default()

	ticketRepo := controllers.New(DB_HOST, DB_PORT)

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
