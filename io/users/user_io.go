package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/users"
)

const usersUrl = api.BASE_URL + "/users"

type Users domain.User

func GetUsers() ([]Users, error) {
	entites := []Users{}
	resp, _ := api.Rest().Get(usersUrl + "/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUser(id string) (Users, error) {
	entity := Users{}
	resp, _ := api.Rest().Get(usersUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUser(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(usersUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUser(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(usersUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUser(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(usersUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
