package models

import "time"

type User struct {
	UserId      string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email,omitempty"`
	Password    string    `json:"password,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	Address     string    `json:"address,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	Birthday    string    `json:"birthday,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Users []User
