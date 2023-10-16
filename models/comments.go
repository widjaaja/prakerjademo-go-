package models

import (
	"gorm.io/gorm"
)


type Comments struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
	Message string `json:"message"`
}