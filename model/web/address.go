package web

type CreateAddressRequest struct {
	City       string `validate:"required" json:"city"`
	Province   string `validate:"required" json:"province"`
	PostalCode int    `validate:"required" json:"postal_code"`
	UserIdFK   int    `validate:"required" json:"user_id_fk"`
}
