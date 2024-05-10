package main

import (
	"belajar-rest-gorm/app"
	"belajar-rest-gorm/controller"
	"belajar-rest-gorm/helper"
	"belajar-rest-gorm/repository"
	"belajar-rest-gorm/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func(cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
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
	r.GET("/user/:id", userController.GetUser)
	r.GET("/users", userController.GetUserList)
	r.PUT("/user/:id", userController.Updateuser)
	r.Logger.Fatal(r.Start(":8080"))
}
