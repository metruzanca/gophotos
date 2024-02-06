package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/metruzanca/go-photos/src/model"
	"github.com/metruzanca/go-photos/src/ui/user"
)

type UserRoute struct{}

func (r UserRoute) UserList(c echo.Context) error {
	u := model.User{
		Email: "s@m.com",
	}
	return render(c, user.List(u))
}
