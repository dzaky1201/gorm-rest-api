package addressservice

import (
	"belajar-rest-gorm/model/entity"
	"belajar-rest-gorm/model/web"
)

type AddressService interface {
	SaveAddress(request web.CreateAddressRequest) (map[string]interface{}, error)
	GetAddressById(addressId int) (entity.AddressEntity, error)
	GetAddressList() ([]entity.AddressEntity, error)
	UpdateAddress(request web.CreateAddressRequest, pathId int) (map[string]interface{}, error)
}
