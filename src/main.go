package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/metruzanca/go-photos/src/routes"
	"github.com/metruzanca/go-photos/src/ui/user"
)

const PORT = "3000"

func main() {
	e := echo.New()
	// app.Use(authMiddleware)
	// e.File("/favicon.png", "/static/favicon.png")
	// e.File("/styles.css", "/static/styles.css")

	e.Static("/", "public")

	userRoute := routes.UserRoute{}
	e.GET("/user", userRoute.UserList)
	e.POST("/clicked", user.HandleClicked)

	fmt.Println("Running on http://localhost:" + PORT)
	e.Start(":" + PORT)
}

// func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// TODO replace with authentication
// 		ctx := context.WithValue(c.Request().Context(), "user", "s@m.dev")
// 		c.SetRequest(c.Request().WithContext(ctx))
// 		return next(c)
// 	}
// }
