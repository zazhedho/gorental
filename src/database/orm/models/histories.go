package models

import "time"

type History struct {
	HistoryId     uint      `gorm:"primaryKey" json:"id"`
	UserId        string    `json:"user_id"`
	VehicleId     uint      `json:"vehicle_id"`
	StartRent     time.Time `json:"start_rent"`
	EndRent       time.Time `json:"end_rent"`
	Prepayment    int       `json:"prepayment"`
	PaymentStatus string    `json:"payment_status"`
	ReturnStatus  string    `json:"return_status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
