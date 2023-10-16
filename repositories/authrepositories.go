package repositories

import (
	"demo-go/configs"
	"demo-go/models"
)

func Register(user *models.User) error {
	result := configs.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}