package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"testDoubleVPartners/database"
	"testDoubleVPartners/models"
)

type TicketRepo struct {
	Database *gorm.DB
}

func New(DB_HOST string, DB_PORT string) *TicketRepo {
	db := database.InitDatabase(DB_HOST, DB_PORT)
	if err := db.AutoMigrate(&models.Ticket{}); err != nil {
		panic(err.Error())
	}

	return &TicketRepo{Database: db}
}

//create ticket
func (repository *TicketRepo) CreateTicket(context *gin.Context) {
	var ticket models.Ticket
	if err := context.BindJSON(&ticket); err != nil {
		panic(err.Error())
	}

	err := models.CreateTicket(repository.Database, &ticket)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, ticket)
}

//get users
func (repository *TicketRepo) GetTickets(context *gin.Context) {
	var ticket []models.Ticket
	err := models.GetTickets(repository.Database, &ticket)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, ticket)
}

//get ticket by id
func (repository *TicketRepo) GetTicket(context *gin.Context) {
	id, _ := context.Params.Get("id")
	var ticket models.Ticket
	err := models.GetTicket(repository.Database, &ticket, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, ticket)
}

// update ticket
func (repository *TicketRepo) UpdateTicket(context *gin.Context) {
	var ticket models.Ticket
	id, _ := context.Params.Get("id")
	err := models.GetTicket(repository.Database, &ticket, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if err := context.BindJSON(&ticket); err != nil {
		panic(err.Error())
	}
	err = models.UpdateTicket(repository.Database, &ticket)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, ticket)
}

// delete ticket
func (repository *TicketRepo) DeleteTicket(context *gin.Context) {
	var ticket models.Ticket
	id, _ := context.Params.Get("id")
	err := models.DeleteTicket(repository.Database, &ticket, id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Ticket deleted successfully",
		"deleted_ticket": ticket,
	})
}
