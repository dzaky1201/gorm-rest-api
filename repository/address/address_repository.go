package address

import "belajar-rest-gorm/model/domain"

type AddressRepository interface {
	CreateAddress(address domain.Address) (domain.Address, error)
	GetAddressById(Id int) (domain.Address, error)
	GetAddresses() ([]domain.Address, error)
	UpdateAddress(address domain.Address) (domain.Address, error)
}
