package address

import (
	"belajar-rest-gorm/model/domain"
	"errors"

	"gorm.io/gorm"
)

type AddressRepositoryImpl struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepositoryImpl {
	return &AddressRepositoryImpl{db: db}
}

func (repo *AddressRepositoryImpl) CreateAddress(address domain.Address) (domain.Address, error) {

	err := repo.db.Create(&address).Error

	if err != nil {
		return domain.Address{}, err
	}

	return address, nil
}

func (repo *AddressRepositoryImpl) GetAddressById(Id int) (domain.Address, error) {
	var addressData domain.Address

	err := repo.db.First(&addressData, "address_id = ?", Id).Error

	if err != nil {
		return domain.Address{}, errors.New("address tidak ditemukan")
	}

	return addressData, nil
}

func (repo *AddressRepositoryImpl) GetAddresses() ([]domain.Address, error) {
	var addresses []domain.Address

	err := repo.db.Find(&addresses).Error

	if err != nil {
		return []domain.Address{}, err
	}

	return addresses, nil
}

func (repo *AddressRepositoryImpl) UpdateAddress(address domain.Address) (domain.Address, error) {
	err := repo.db.Model(domain.Address{}).Where("address_id = ?", address.AddressID).Updates(address).Error

	if err != nil {
		return address, err
	}

	return address, nil
}
