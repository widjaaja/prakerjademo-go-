package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
	Password string `gorm:"not null" json:"password"`
}

type UserResponse struct {
	gorm.Model
	Name string `json:"name"`
	Token string `json:"token"`
}