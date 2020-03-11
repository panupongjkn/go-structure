package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/panupong25509/go-structure/actions/repositories"
)

func GetUser(c echo.Context) error {
	user, err := repositories.GetUser(c)
	if err != nil {
		return c.JSON(404, "")
	}
	return c.JSON(http.StatusOK, user)
}
