package login

import (
	"errors"
	"fmt"
	"obas/src/api"
	loginDomain "obas/src/domain/login"
	domain "obas/src/domain/users"
)

const loginURL = api.BASE_URL + "/login"

type Register loginDomain.Register
type Login loginDomain.Login
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

func Login_io(email, password string) ([]Token, error) {
	entity := Login{email, password}
	entites := []Token{}

	resp, _ := api.Rest().SetBody(entity).Post(loginURL + "/login")
	//result1,_:=json.Marshal(resp)
	//fmt.Printf("%s\n",result1)
	result := api.JSON.Unmarshal(resp.Body(), &entites)

	fmt.Printf("%s\n", result)

	println("the entities is:", result)

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	return entites, nil

}
