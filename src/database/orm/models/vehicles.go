package models

import "time"

type Vehicle struct {
	VehicleId   uint      `gorm:"primaryKey" json:"id,omitempty"`
	Name        string    `gorm:"not null" json:"name,omitempty"`
	Location    string    `gorm:"not null" json:"location,omitempty"`
	Description string    `gorm:"not null" json:"description,omitempty"`
	Price       int       `gorm:"not null" json:"price,omitempty"`
	Status      string    `gorm:"not null" json:"status,omitempty"`
	Stock       int       `gorm:"not null" json:"stock,omitempty"`
	Category    string    `gorm:"not null" json:"category,omitempty"`
	Rating      float32   `gorm:"not null" json:"rating,omitempty"`
	Image       string    `gorm:"not null" json:"image,omitempty"`
	CreatedAt   time.Time `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:now(); not null" json:"updated_at"`
}

type Vehicles []Vehicle
