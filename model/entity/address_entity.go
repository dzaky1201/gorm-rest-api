package entity

import "belajar-rest-gorm/model/domain"

type AddressEntity struct {
	ID         int    `json:"address_id"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode int    `json:"postal_code"`
}

func ToAddressEntity(id int, city string, province string, postalCode int) AddressEntity {
	return AddressEntity{
		ID:         id,
		City:       city,
		Province:   province,
		PostalCode: postalCode,
	}
}

func ToAddressListEntity(addresses []domain.Address) []AddressEntity {
	addressData := []AddressEntity{}

	for _, address := range addresses {
		addressData = append(addressData, ToAddressEntity(address.AddressID, address.City, address.Province, address.PostalCode))
	}

	return addressData
}
