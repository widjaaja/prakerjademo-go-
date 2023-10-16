package repositories

import (
	"demo-go/configs"
	"demo-go/models"
)

func GetNews(newsList *[]models.News) error {
	result := configs.DB.Find(&newsList)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddNews(newsDB *models.News) error {
	result := configs.DB.Create(&newsDB)
	if result.Error != nil {
		return result.Error
	}
	return nil
}