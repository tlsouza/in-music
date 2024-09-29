package types

import "time"

// Profile struct representing a user profile
type Profile struct {
	ID        uint64 `json:"id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Product struct {
	SKU *string `json:"SKU"`
}

//Http requests structures
type ProfileHttpRequest struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type ProfileHttpResponse struct {
	ID                   uint64                       `json:"id"`
	Email                string                       `json:"email"`
	Firstname            string                       `json:"firstname"`
	Lastname             string                       `json:"lastname"`
	ProductRegistrations []ProductRegistrationHttpRes `json:"product_registrations"`
}

type ProductRegistrationHttpReq struct {
	PurchaseDate                   time.Time                    `json:"purchase_date"`
	ExpiryAt                       *time.Time                   `json:"expiry_at"`
	Product                        Product                      `json:"product"`
	SerialCode                     string                       `json:"serial_code"`
	AdditionalProductRegistrations []ProductRegistrationHttpReq `json:"additional_product_registrations"`
}

type ProductRegistrationHttpResChild struct {
	Id           uint64     `json:"id"`
	PurchaseDate time.Time  `json:"purchase_date"`
	ExpiryAt     *time.Time `json:"expiry_at"`
	Product      Product    `json:"product"`
	SerialCode   string     `json:"serial_code"`
}

type ProductRegistrationHttpRes struct {
	Id                             uint64                            `json:"id"`
	PurchaseDate                   time.Time                         `json:"purchase_date"`
	ExpiryAt                       *time.Time                        `json:"expiry_at"`
	Product                        Product                           `json:"product"`
	SerialCode                     string                            `json:"serial_code"`
	AdditionalProductRegistrations []ProductRegistrationHttpResChild `json:"additional_product_registrations"`
}

type ProductRegistration struct {
	Id           uint64     `json:"id"`
	PurchaseDate time.Time  `json:"purchase_date"`
	ExpiryAt     *time.Time `json:"expiry_at"`
	Product      Product    `json:"product"`
	SerialCode   string     `json:"serial_code"`
	ProfileId    *uint64    `json:"profile_id"`
	ParentId     *uint64    `json:"parent_id"`
	RootId       *uint64    `json:"root_id"`
}
