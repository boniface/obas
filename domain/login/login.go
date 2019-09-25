package domain

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Email string `json:"email"`
}
