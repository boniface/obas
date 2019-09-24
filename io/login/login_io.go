package login

import (
	"encoding/json"
	"errors"
	"net/mail"
	"obas/api"
	"obas/controllers/login"
	"obas/src/domain/users"
)

const roleurl = api.BASE_URL + "/login"

type User users.User
type MessageResponse mail.Message
type UserGeneratedToken security.UserGeneratedToken
type Account users.Account
type Profile login.Profile
type LoginCredential login.LoginCredential

func ResetPasswordRequest(resetKey string) (MessageResponse, error) {
	entity := MessageResponse{}
	resp, _ := api.Rest().Get(roleurl + "/reset/" + resetKey)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}

func IsUserRegistered(user User) (bool, error) {
	resp, _ := api.Rest().
		SetBody(user).
		Post(roleurl + "/registered")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil

}

func ForgotPassword(profile Profile) (MessageResponse, error) {
	entity := MessageResponse{}
	resp, _ := api.Rest().
		SetBody(profile).
		Post(roleurl + "/forgotpassword")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}

func GetUserAccounts(email string) ([]Account, error) {
	entities := []Account{}
	resp, _ := api.Rest().Get(roleurl + "/user/" + email)
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil

}

func GetLoginToken(loginCredential LoginCredential) (UserGeneratedToken, error) {
	entity := UserGeneratedToken{}
	resp, _ := api.Rest().
		SetBody(loginCredential).
		Post(roleurl + "/getlogintoken")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}

func IsUserLoggedIn(user User) (bool, error) {
	// GET THIS FROM SESSION
	//resp, _ := api.Rest().
	//	SetBody(entity).
	//	Post(roleurl + "/create")
	//if resp.IsError() {
	//	return false, errors.New(resp.Status())
	//}

	return true, nil

}

func GetUserRole(user User) (string, error) {
	// GET THIS FROM THE SESSION

	//role := Role{}
	//resp, _ := api.Rest().Get(roleurl + "/get/" + id)
	//if resp.IsError() {
	//	return role , errors.New(resp.Status())
	//}
	//err := json.Unmarshal(resp.Body(), &role)
	//if err != nil {
	//	return role ,  errors.New(resp.Status())
	//}
	return "", nil

}

func Logout() (bool, error) {
	// INVALIDATE THE SESSION

	//resp, _ := api.Rest().
	//	SetBody(entity).
	//	Post(roleurl + "/create")
	//if resp.IsError() {
	//	return false, errors.New(resp.Status())
	//}

	return true, nil

}
