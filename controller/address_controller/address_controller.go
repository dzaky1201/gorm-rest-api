package addresscontroller

import "github.com/labstack/echo/v4"

type AddressController interface {
	SaveAddress(c echo.Context) error
	GetAddress(c echo.Context) error
	GetAddressList(c echo.Context) error
	UpdateAddress(c echo.Context) error
}
