package login

import (
	"errors"
	"obas/src/api"
	loginDomain "obas/src/domain/login"
)

const loginURL = api.BASE_URL + "/login"

type Register loginDomain.Register

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
