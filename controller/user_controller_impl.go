package controller

import (
	"belajar-rest-gorm/model"
	"belajar-rest-gorm/model/web"
	"belajar-rest-gorm/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		userService: service,
	}
}

func (controller *UserControllerImpl) SaveUser(c echo.Context) error {
	user := new(web.UserServiceRequest)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	saveUser, errSaveUser := controller.userService.SaveUser(*user)

	if errSaveUser != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveUser.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "berhasil membuat user", saveUser))
}

func (controller *UserControllerImpl) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	getUser, errGetUser := controller.userService.GetUser(id)

	if errGetUser != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, errGetUser.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", getUser))
}
