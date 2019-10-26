package domain

import "time"

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Email string `json:"email"`
}

type ForgetPassword struct {
	Email string `json:"email"`
}

type ResetPassword struct {
	Password string `json:"password"`
}

type LoginToken struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type ChangePassword struct {
	Email           string    `json:"email"`
	OldPassword     string    `json:"oldPassword"`
	NewPassword     string    `json:"newPassword"`
	DatetimeChanged time.Time `json:"datetimeChanged"`
}
