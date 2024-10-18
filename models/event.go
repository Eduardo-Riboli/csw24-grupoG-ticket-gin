package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model         
	TenantID    uint   `gorm:"not null"`              
	Name   		string `gorm:"size:100;not null"`      
	Type        string `gorm:"size:50;not null"`       
	Location    string `gorm:"size:255"`               
	Date     	time.Time `gorm:"not null"`          // Date and time of the event
	Tickets     []Ticket `gorm:"foreignKey:EventID"` // One-to-many relationship with Ticket
}

