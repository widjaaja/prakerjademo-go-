package main

import (

	"demo-go/configs"
	"demo-go/routes"

	"github.com/labstack/echo/v4"
)

func main(){
	configs.LoadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(":8000")
}






