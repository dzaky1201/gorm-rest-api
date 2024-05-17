package app

import (
	"belajar-rest-gorm/controller"
	addresscontroller "belajar-rest-gorm/controller/address_controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(user controller.UserController, address addresscontroller.AddressController)(e *echo.Echo){

}