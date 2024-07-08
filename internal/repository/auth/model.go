package auth

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password []byte `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
