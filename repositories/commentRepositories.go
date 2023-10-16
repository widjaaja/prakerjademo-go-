package repositories

import (
    "demo-go/configs"
    "demo-go/models"
)

func GetComments(commentLists *[]models.Comments) error {
	result := configs.DB.Find(&commentLists)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddComments(commentsDB *models.Comments) error {
	result := configs.DB.Create(&commentsDB)
	if result.Error != nil {
		return result.Error
	}
	return nil
}