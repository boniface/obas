package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/users"
)

const usersComUrl = api.BASE_URL + "/users"

type UsersCom domain.UserCommunication

func GetUserCommunications() ([]UsersCom, error) {
	entites := []UsersCom{}
	resp, _ := api.Rest().Get(usersComUrl + "/communication/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUserCommunication(id string) (UsersCom, error) {
	entity := UsersCom{}
	resp, _ := api.Rest().Get(usersComUrl + "/communication/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserCommunication(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(usersComUrl + "/communication/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUserCommunication(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(usersComUrl + "/communication/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUserCommunication(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(usersComUrl + "/communication/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
