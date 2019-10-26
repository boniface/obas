package login

import (
	"encoding/json"
	"errors"
	"obas/api"
	loginDomain "obas/domain/login"
)

const loginURL = api.BASE_URL + "/login"

type Register loginDomain.Register
type Forget loginDomain.ForgetPassword
type Login loginDomain.Login
type Password loginDomain.ResetPassword
type LoginToken loginDomain.LoginToken
type ChangePassword loginDomain.ChangePassword

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

func DoForgetPassword(email string) (bool, error) {
	entity := Forget{}
	entity.Email = email
	resp, _ := api.Rest().
		SetBody(entity).
		Post(loginURL + "/forgotpassword")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func DoReset(resetKey string) (bool, error) {
	resp, _ := api.Rest().Get(loginURL + "/passwordreset/" + resetKey)
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
	//respEntity = LoginToken{entity.Email, "aerefasd.foqerwfdasdfaoduo"}
	err := json.Unmarshal(resp.Body(), &respEntity)
	if err != nil {
		return respEntity, errors.New(resp.Status())
	}
	return respEntity, nil
}

func DoChangePassword(entity ChangePassword, token string) (LoginToken, error) {
	resp, _ := api.Rest().
		SetAuthToken(token).
		SetBody(entity).
		Post(loginURL + "/changepassword")
	if resp.IsError() {
		return LoginToken{}, errors.New(resp.Status())
	}
	respEntity := LoginToken{}
	//respEntity = LoginToken{entity.Email, "fdgsfjuipopoqqg$&aerefasd.dafadfwe4534ff0ygyj55r"}
	err := json.Unmarshal(resp.Body(), &respEntity)
	if err != nil {
		return respEntity, errors.New(resp.Status())
	}
	return respEntity, nil
}
