package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	SaveUser(c echo.Context) error
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
}
