package login

import (
	"errors"
	"fmt"
	"obas/src/api"
	loginDomain "obas/src/domain/login"
)

const loginURL = api.BASE_URL + "/login"

type Register loginDomain.Register
type Login loginDomain.Login

func DoRegister(email string) (bool, error) {
	entity := Register{}
	entity.Email = email
	resp, _ := api.Rest().
		SetBody(entity).
		Post(loginURL + "/register")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
func Login_io(email, password string) (bool, error) {
	entity := Login{email, password}

	resp, _ := api.Rest().SetBody(entity).Post(loginURL + "/login")

	fmt.Print(" The Result ")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
