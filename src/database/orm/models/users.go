package models

import "time"

type User struct {
	UserId      string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	DisplayName string    `json:"display_name"`
	Birthday    string    `json:"birthday"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Users []User
