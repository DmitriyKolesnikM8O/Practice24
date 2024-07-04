package user

// User
type Product struct {
	ID           string `json:"id" bson:"_id, omitempty"`
	Email        string `json:"email" bson:"email"`
	Name         string `json:"name" bson:"name"`
	PasswordHash string `json:"-" bson:"password"`
}

// CreateUserTDO
type CreateProductDTO struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
