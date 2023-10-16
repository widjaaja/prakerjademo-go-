package repositories

import (
	"errors"
	"fmt"

	"demo-go/configs"
	"demo-go/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/labstack/echo/v4"
)

func Register(user *models.User) error {
	result := configs.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Login(c echo.Context, user *models.User) error {
	// Find the user by username
	var dbUser models.User
	result := configs.DB.Where("name = ?", user.Name).First(&dbUser)
	if result.Error != nil {
		fmt.Println(result.Error, "Invalid password")
        return result.Error
	}

	// Compare the provided password with the hashed password
	pass := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if pass != nil {
		err := errors.New("Invalid password")
		fmt.Println(err, "Invalid password")
        return err
	}
	return nil
}