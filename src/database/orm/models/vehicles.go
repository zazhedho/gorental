package models

import "time"

type Vehicle struct {
	VehicleId   uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Location    string    `json:"location,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       int       `json:"price,omitempty"`
	Status      string    `json:"status,omitempty"`
	Stock       int       `json:"stock,omitempty"`
	Category    string    `json:"category,omitempty"`
	Image       string    `json:"image,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Rating      float32   `json:"rating,omitempty"`
}

type Vehicles []Vehicle
