package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UserDTO struct {
	Nickname         *string `json:"nickname"`
	CurrentPassword  *string `json:"currentPassword"`
	Password         *string `json:"password" validate:"min:8"`
	Email            *string    `gorm:"size:100;not null;unique" json:"email"`
}

type UserResponse struct {
	ID             uint32       `json:"id"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	NickName       string     `json:"nickname"`
	Email          string     `json:"email"`
}