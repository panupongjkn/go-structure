package routes

import (
	"github.com/labstack/echo"
	"github.com/panupong25509/go-structure/actions/handlers"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/user", handlers.GetUser)
	return e
}
