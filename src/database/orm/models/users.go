package models

import "time"

type User struct {
	UserId      string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id,omitempty"`
	Username    string    `json:"username,omitempty" validate:"required"`
	Email       string    `json:"email,omitempty" validate:"required"`
	Password    string    `json:"password,omitempty" validate:"required"`
	Phone       string    `json:"phone,omitempty" validate:"required"`
	Gender      string    `json:"gender,omitempty"`
	Address     string    `json:"address,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	Birthday    string    `json:"birthday,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Role        string    `json:"role,omitempty"`
}

type Users []User
