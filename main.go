package main

import (
	"belajar-rest-gorm/app"
	"belajar-rest-gorm/controller"
	"belajar-rest-gorm/helper"
	"belajar-rest-gorm/repository"
	"belajar-rest-gorm/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
func main() {
	db := app.DBConnection()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := echo.New()
	r.Validator = &CustomValidator{validator: validator.New()}
	r.HTTPErrorHandler = helper.BindAndValidate
	r.POST("/register", userController.SaveUser)
	r.Logger.Fatal(r.Start(":8080"))
}
