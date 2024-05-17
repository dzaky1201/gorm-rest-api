package main

import (
	"belajar-rest-gorm/app"
	"belajar-rest-gorm/controller"
	addresscontroller "belajar-rest-gorm/controller/address_controller"
	"belajar-rest-gorm/helper"
	"belajar-rest-gorm/model"
	"belajar-rest-gorm/repository"
	"belajar-rest-gorm/repository/address"
	"belajar-rest-gorm/service"
	addressservice "belajar-rest-gorm/service/address_service"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := app.DBConnection()
	userRepo := repository.NewUserRepository(db)
	tokenUseCase := helper.NewTokenUseCase()
	userService := service.NewUserService(userRepo, tokenUseCase)
	userController := controller.NewUserController(userService)

	addressRepo := address.NewAddressRepository(db)
	addressService := addressservice.NewAddressService(addressRepo)
	addressController := addresscontroller.NewAddressController(addressService)

	r := echo.New()
	r.Validator = CustomValidator{validator: validator.New()}
	r.HTTPErrorHandler = helper.BindAndValidate
	r.POST("/register", userController.SaveUser)
	r.GET("/user/:id", userController.GetUser, JWTProtection())
	r.GET("/users", userController.GetUserList, JWTProtection())
	r.PUT("/user/:id", userController.Updateuser, JWTProtection())
	r.POST("/user/login", userController.LoginUser)

	r.POST("/address/add", addressController.SaveAddress)
	r.GET("/address/:id", addressController.GetAddress, JWTProtection())
	r.GET("/addresses", addressController.GetAddressList, JWTProtection())
	r.PUT("/address/:id", addressController.UpdateAddress, JWTProtection())
	r.Logger.Fatal(r.Start(":8080"))
}

func JWTProtection() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, "anda harus login untuk mengakses resource ini", nil))
		},
	})
}
