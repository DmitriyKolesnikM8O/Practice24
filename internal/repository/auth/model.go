package auth

// TODO HashPassword
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password []byte `json:"password"`
}

type UserRegister struct {
	ID       int    `json:"-"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
