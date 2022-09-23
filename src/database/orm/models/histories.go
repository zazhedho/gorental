package models

import "time"

type History struct {
	HistoryId     uint      `gorm:"primaryKey" json:"history_id"`
	UserId        string    `gorm:"not null" json:"user_id"`
	User          User      `gorm:"not null" json:"user_data"`
	VehicleId     uint      `gorm:"not null" json:"vehicle_id"`
	Vehicle       Vehicle   `gorm:"not null" json:"vehicle_data"`
	StartRent     string    `gorm:"not null" json:"start_rent"`
	EndRent       string    `gorm:"not null" json:"end_rent"`
	Prepayment    int       `gorm:"not null" json:"prepayment"`
	PaymentStatus string    `gorm:"not null" json:"payment_status"`
	ReturnStatus  string    `gorm:"not null" json:"return_status"`
	CreatedAt     time.Time `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:now(); not null" json:"updated_at"`
}

type Histories []History
