package users

import (
	"errors"
	"fmt"
	"obas/api"
	domain "obas/domain/users"
)

const usersPwdUrl = api.BASE_URL + "/users/password/"

type UserPassword domain.UserPassword

func GetUserPasswords() ([]UserPassword, error) {
	entites := []UserPassword{}
	resp, serverEr := api.Rest().Get(usersPwdUrl + "all")

	if resp.IsError() {
		fmt.Println(" Is request from Server Okay", serverEr)
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUserPassword(id string) (UserPassword, error) {
	entity := UserPassword{}
	resp, _ := api.Rest().Get(usersPwdUrl + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserPassword(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(usersPwdUrl + "create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func UpdateUserPassword(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(usersPwdUrl + "update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func DeleteUserPassword(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(usersPwdUrl + "delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
