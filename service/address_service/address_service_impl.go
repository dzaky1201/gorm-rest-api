package addressservice

import (
	"belajar-rest-gorm/helper"
	"belajar-rest-gorm/model/domain"
	"belajar-rest-gorm/model/entity"
	"belajar-rest-gorm/model/web"
	"belajar-rest-gorm/repository/address"
)

type AddressServiceImpl struct {
	repository address.AddressRepository
}

func NewAddressService(repository address.AddressRepository) *AddressServiceImpl {
	return &AddressServiceImpl{
		repository: repository,
	}
}

func (service *AddressServiceImpl) SaveAddress(request web.CreateAddressRequest) (map[string]interface{}, error) {

	addressReq := domain.Address{
		City:       request.City,
		Province:   request.Province,
		UserIDFK:   request.UserIdFK,
		PostalCode: request.PostalCode,
	}

	saveAddress, errSaveAddress := service.repository.CreateAddress(addressReq)

	if errSaveAddress != nil {
		return nil, errSaveAddress
	}

	return helper.ResponseToJson{"city": saveAddress.City, "province": saveAddress.Province, "postal_code": saveAddress.PostalCode}, nil

}

func (service *AddressServiceImpl) GetAddressById(addressId int) (entity.AddressEntity, error) {
	getAddress, errGetAddress := service.repository.GetAddressById(addressId)

	if errGetAddress != nil {
		return entity.AddressEntity{}, errGetAddress
	}

	return entity.ToAddressEntity(getAddress.AddressID, getAddress.City, getAddress.Province, getAddress.PostalCode), nil
}

func (service *AddressServiceImpl) GetAddressList() ([]entity.AddressEntity, error) {
	getAddressList, errGetAddressList := service.repository.GetAddresses()

	if errGetAddressList != nil {
		return []entity.AddressEntity{}, errGetAddressList
	}

	return entity.ToAddressListEntity(getAddressList), nil
}

func (service *AddressServiceImpl) UpdateAddress(request web.CreateAddressRequest, pathId int) (map[string]interface{}, error) {
	getAddressrById, err := service.repository.GetAddressById(pathId)
	if err != nil {
		return nil, err
	}

	if request.City == "" {
		request.City = getAddressrById.City
	}

	if request.PostalCode == 0 {
		request.PostalCode = getAddressrById.PostalCode
	}

	if request.Province == "" {
		request.Province = getAddressrById.Province
	}

	if request.UserIdFK == 0 {
		request.UserIdFK = getAddressrById.UserIDFK
	}

	addressRequest := domain.Address{
		AddressID:  getAddressrById.AddressID,
		City:       request.City,
		PostalCode: request.PostalCode,
		Province: request.Province,
		UserIDFK:   request.UserIdFK,
	}

	updateAddress, errUpdate := service.repository.UpdateAddress(addressRequest)

	if errUpdate != nil {
		return nil, errUpdate
	}

	return helper.ResponseToJson{"city": updateAddress.City, "province": updateAddress.Province, "postal_code": updateAddress.PostalCode}, nil
}
