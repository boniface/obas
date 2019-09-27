package login

import (
	"encoding/json"
	"errors"
	"obas/api"
	loginDomain "obas/domain/login"
)

const loginURL = api.BASE_URL + "/login"

type Register loginDomain.Register
type Login loginDomain.Login
type LoginToken loginDomain.LoginToken

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

func DoLogin(email string, password string) (LoginToken, error) {
	entity := Login{email, password}
	resp, _ := api.Rest().
		SetBody(entity).
		Post(loginURL + "/login")
	if resp.IsError() {
		return LoginToken{}, errors.New(resp.Status())
	}
	respEntity := LoginToken{}
	err := json.Unmarshal(resp.Body(), &respEntity)
	if err != nil {
		return respEntity, errors.New(resp.Status())
	}
	return respEntity, nil
}
