package controllers

import (
	"net/http"
	"demo-go/middlewares"
	"demo-go/models"
	"demo-go/repositories"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(c echo.Context) error {
	var userRegister models.User
	c.Bind(&userRegister)

	result, _ := bcrypt.GenerateFromPassword(
		[]byte(userRegister.Password),
        bcrypt.DefaultCost,
	)
	userRegister.Password = string(result)
	err := repositories.Register(&userRegister)
	
	if err != nil {
        return c.JSON(http.StatusInternalServerError, models.BaseResponse{
		    Message: err.Error(),
            Status: false,
            Data: nil,
		})
    }

	var userResponse models.UserResponse
	userResponse.ID = userRegister.ID
	userResponse.Name = userRegister.Name
	userResponse.Token = middlewares.GenerateJWTToken(
		userResponse.ID,
		userResponse.Name,
	)

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil registered successfully",
        Status: true,
        Data: userResponse,
	})
}