package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/metruzanca/go-photos/ui/user"
)

type UserRoute struct{}

func (r UserRoute) UserList(c echo.Context) error {
	return render(c, user.List())
}
