package models

import "time"

type Vehicle struct {
	VehicleId   uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Status      string    `json:"status"`
	Stock       int       `json:"stock"`
	Category    string    `json:"category"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Vehicles []Vehicle
