// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"belajar-rest-gorm/app"
	"belajar-rest-gorm/controller"
	"belajar-rest-gorm/controller/address_controller"
	"belajar-rest-gorm/helper"
	"belajar-rest-gorm/repository"
	"belajar-rest-gorm/repository/address"
	"belajar-rest-gorm/service"
	"belajar-rest-gorm/service/address_service"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

// Injectors from injector.go:

func StartServer() *echo.Echo {
	db := app.DBConnection()
	userRepositoryImpl := repository.NewUserRepository(db)
	tokenUseCaseImpl := helper.NewTokenUseCase()
	userServiceImpl := service.NewUserService(userRepositoryImpl, tokenUseCaseImpl)
	userControllerImpl := controller.NewUserController(userServiceImpl)
	addressRepositoryImpl := address.NewAddressRepository(db)
	addressServiceImpl := addressservice.NewAddressService(addressRepositoryImpl)
	addressControllerImpl := addresscontroller.NewAddressController(addressServiceImpl)
	echoEcho := app.IntializedServe(userControllerImpl, addressControllerImpl)
	return echoEcho
}

// injector.go:

var userSet = wire.NewSet(repository.NewUserRepository, wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)), helper.NewTokenUseCase, wire.Bind(new(helper.TokenUseCase), new(*helper.TokenUseCaseImpl)), service.NewUserService, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)), controller.NewUserController, wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)))

var addressSet = wire.NewSet(address.NewAddressRepository, wire.Bind(new(address.AddressRepository), new(*address.AddressRepositoryImpl)), addressservice.NewAddressService, wire.Bind(new(addressservice.AddressService), new(*addressservice.AddressServiceImpl)), addresscontroller.NewAddressController, wire.Bind(new(addresscontroller.AddressController), new(*addresscontroller.AddressControllerImpl)))
