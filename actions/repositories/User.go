package repositories

import (
	"github.com/labstack/echo"
	"github.com/panupong25509/go-structure/models"
)

func GetUser(c echo.Context) (interface{}, error) {
	user := models.User{}
	user.CreateUser("Panupong")
	return user, nil
}
