package auth

type LoginDao struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}