package main

import (
	"belajar-rest-gorm/controller"
	addresscontroller "belajar-rest-gorm/controller/address_controller"
	"belajar-rest-gorm/repository"
	"belajar-rest-gorm/repository/address"
	"belajar-rest-gorm/service"
	addressservice "belajar-rest-gorm/service/address_service"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var userSet = wire.NewSet(
	repository.NewUserRepository,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	service.NewUserService,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
	controller.NewUserController,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var addressSet = wire.NewSet(
	address.NewAddressRepository,
	wire.Bind(new(address.AddressRepository), new(*address.AddressRepositoryImpl)),
	addressservice.NewAddressService,
	wire.Bind(new(addressservice.AddressService), new(*addressservice.AddressServiceImpl)),
	addresscontroller.NewAddressController,
	wire.Bind(new(addresscontroller.AddressController), new(*addresscontroller.AddressControllerImpl)),
)

func RunServer() (e *echo.Echo) {
	return
}
