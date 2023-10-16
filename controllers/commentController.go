package controllers

import (
    "net/http"
    "demo-go/models"
    "demo-go/repositories"

	"github.com/labstack/echo/v4"
)

func GetCommentController(c echo.Context) error {
	var comments []models.Comments
	
	err := repositories.GetComments(&comments)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
            Status: false,
            Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "success",
        Status: true,
        Data: comments,
	})
}

func CreateCommentsController(c echo.Context) error {
	var commentsRequest models.Comments
	c.Bind(&commentsRequest)

	err := repositories.AddComments(&commentsRequest)

	if err != nil {
        return c.JSON(http.StatusInternalServerError, models.BaseResponse{
		    Message: err.Error(),
            Status: false,
            Data: nil,
		})
    }
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menambah data",
        Status: true,
        Data: commentsRequest,
	})
} 