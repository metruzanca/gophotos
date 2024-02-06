package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/metruzanca/go-photos/src/routes"
)

const PORT = "3000"

func main() {
	app := echo.New()
	app.Use(authMiddleware)

	userRoute := routes.UserRoute{}
	app.GET("/user", userRoute.UserList)

	app.Start(":" + PORT)
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO replace with authentication
		ctx := context.WithValue(c.Request().Context(), "user", "s@m.dev")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
