package models

import (
	"gorm.io/gorm"
	"time"
)

type Ticket struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	User string `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status bool `json:"status"`
}


func CreateTicket(database *gorm.DB, Ticket *Ticket) (err error) {
   err = database.Create(Ticket).Error
   if err != nil {
      return err
   }
   return nil
}


func GetTickets(database *gorm.DB, Ticket *[]Ticket) (err error) {
   err = database.Find(Ticket).Error
   if err != nil {
      return err
   }
   return nil
}


func GetTicket(database *gorm.DB, Ticket *Ticket, id string) (err error) {
   err = database.Where("id = ?", id).First(Ticket).Error
   if err != nil {
      return err
   }
   return nil
}


func UpdateTicket(database *gorm.DB, Ticket *Ticket) (err error) {
   database.Save(Ticket)
   return nil
}


func DeleteTicket(database *gorm.DB, Ticket *Ticket, id string) (err error) {
   database.Where("id = ?", id).Delete(Ticket)
   return nil
}
