package domain

type Login struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type ForgottenPassword struct {
	Email string `json:"Email"`
}
