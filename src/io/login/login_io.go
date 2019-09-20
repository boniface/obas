package login

import (
	"errors"
	"obas/src/api"
	loginDomain "obas/src/domain/login"
	domain "obas/src/domain/users"
)

const loginURL = api.BASE_URL + "/login"

type Register loginDomain.Register
type User domain.User
type Token loginDomain.LoginToken

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

func IsUserRegistered(email string) (bool, error) {
	entity := User{}
	entity.Email = email
	resp, _ := api.Rest().
		SetBody(entity).
		Post(loginURL + "/registered")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func GetLoginToken(userToken loginDomain.LoginToken) {

}
