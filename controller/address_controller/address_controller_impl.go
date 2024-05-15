package addresscontroller

import (
	"belajar-rest-gorm/model"
	"belajar-rest-gorm/model/web"
	addressservice "belajar-rest-gorm/service/address_service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AddressControllerImpl struct {
	service addressservice.AddressService
}

func NewAddressController(service addressservice.AddressService) *AddressControllerImpl {
	return &AddressControllerImpl{
		service: service,
	}
}

func (controller *AddressControllerImpl) SaveAddress(c echo.Context) error {
	address := new(web.CreateAddressRequest)

	if err := c.Bind(address); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(address); err != nil {
		return err
	}

	saveAddress, errSaveAddress := controller.service.SaveAddress(*address)

	if errSaveAddress != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveAddress.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "berhasil membuat address", saveAddress))
}

func (controller *AddressControllerImpl) GetAddress(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	getAddress, errGetAddress := controller.service.GetAddressById(id)

	if errGetAddress != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, errGetAddress.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", getAddress))
}

func (controller *AddressControllerImpl) GetAddressList(c echo.Context) error {
	getUsers, errGetUsers := controller.service.GetAddressList()

	if errGetUsers != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, errGetUsers.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", getUsers))
}

func (controller *AddressControllerImpl) UpdateAddress(c echo.Context) error {

	address := new(web.CreateAddressRequest)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(address); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	addressUpdate, errAddressUpdate := controller.service.UpdateAddress(*address, id)

	if errAddressUpdate != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errAddressUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "data berhasil diupdate", addressUpdate))
}
