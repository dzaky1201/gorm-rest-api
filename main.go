package main

import (
	"belajar-rest-gorm/app"
	"belajar-rest-gorm/controller"
	"belajar-rest-gorm/repository"
	"belajar-rest-gorm/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	db := app.DBConnection()
	userRepo := repository.NewUserRepository(db)
	validate := validator.New()
	userService := service.NewUserService(userRepo, validate)
	userController := controller.NewUserController(userService)

	r := echo.New()
	r.Debug = true
	r.POST("/register", userController.SaveUser)
	r.Logger.Fatal(r.Start(":8080"))
}