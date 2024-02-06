package main

import (
	"github.com/labstack/echo/v4"
	"github.com/metruzanca/go-photos/routes"
)

const PORT = "3000"

func main() {
	app := echo.New()

	userRoute := routes.UserRoute{}
	app.GET("/user", userRoute.UserList)

	app.Start(":" + PORT)
}
