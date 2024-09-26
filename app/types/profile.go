package types

// Profile struct representing a user profile
type Profile struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Profile struct representing a user profile
type ProfileHttpRequest struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Product struct {
	SKU string `json:"SKU"`
}
