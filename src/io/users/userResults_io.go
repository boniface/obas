package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/users"
)

const userResultUrl = api.BASE_URL + "/users"

type uResults domain.UserResults

func GetUserResults() ([]uResults, error) {
	entites := []uResults{}
	resp, _ := api.Rest().Get(userResultUrl + "/results/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUserResult(id string) (uResults, error) {
	entity := uResults{}
	resp, _ := api.Rest().Get(userResultUrl + "/results/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserResults(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userResultUrl + "/results/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUserResults(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userResultUrl + "/results/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUserResults(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userResultUrl + "/results/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
