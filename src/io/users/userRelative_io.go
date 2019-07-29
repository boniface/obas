package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/users"
)

const userRelUrl = api.BASE_URL + "/relative/users"

type uRelative domain.UserRelative

func GetUserRelatives() ([]uRelative, error) {
	entites := []uRelative{}
	resp, _ := api.Rest().Get(userRelUrl + "/relative/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUserRelative(id string) (uRelative, error) {
	entity := uRelative{}
	resp, _ := api.Rest().Get(userRelUrl + "/relative/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserRelative(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userRelUrl + "/relative/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUserRelative(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userRelUrl + "/relative/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUserRelative(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userRelUrl + "/relative/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
