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

type ProductRegistrationHttpRequestChild struct {
	PurchaseDate time.Time  `json:"purchase_date"`
	ExpiryAt     *time.Time `json:"expiry_at"`
	Product      Product    `json:"product"`
	SerialCode   string     `json:"serial_code"`
}
type ProductRegistrationHttpRequest struct {
	PurchaseDate                   time.Time                             `json:"purchase_date"`
	ExpiryAt                       *time.Time                            `json:"expiry_at"`
	Product                        Product                               `json:"product"`
	SerialCode                     string                                `json:"serial_code"`
	ProfileId                      uint64                                `json:"profile_id"`
	AdditionalProductRegistrations []ProductRegistrationHttpRequestChild `json:"additional_product_registrations"`
}
