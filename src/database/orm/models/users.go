package models

import "time"

type User struct {
	UserId      string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id,omitempty"`
	Username    string    `gorm:"not null" json:"username,omitempty" validate:"required"`
	Email       string    `gorm:"not null" json:"email,omitempty" validate:"required"`
	Password    string    `gorm:"not null" json:"password,omitempty" validate:"required"`
	Phone       string    `gorm:"not null" json:"phone,omitempty" validate:"required"`
	Gender      string    `gorm:"not null" json:"gender,omitempty"`
	Address     string    `gorm:"not null" json:"address,omitempty"`
	DisplayName string    `gorm:"not null" json:"display_name,omitempty"`
	Birthday    string    `gorm:"not null" json:"birthday,omitempty"`
	Role        string    `gorm:"not null" json:"role,omitempty"`
	CreatedAt   time.Time `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:now(); not null" json:"updated_at"`
}

type Users []User
