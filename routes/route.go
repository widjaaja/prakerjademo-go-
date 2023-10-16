package routes

import (
	"os"
	"demo-go/controllers"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.Use(middleware.Logger())
	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	eAuth.POST("/comments", controllers.CreateCommentsController)
	eAuth.GET("/comments", controllers.GetCommentController)
	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)
}