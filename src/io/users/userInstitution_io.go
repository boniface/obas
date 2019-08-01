package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/users"
)

const userIntUrl = api.BASE_URL + "/users"

type UsersInt domain.UserInstitution

func GetUserIntstitutions() ([]UsersInt, error) {
	entites := []UsersInt{}
	resp, _ := api.Rest().Get(userIntUrl + "/institution/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUserIntstitution(id string) (UsersInt, error) {
	entity := UsersInt{}
	resp, _ := api.Rest().Get(userIntUrl + "/institution/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserIntstitution(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userIntUrl + "/institution/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUserIntstitution(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userIntUrl + "/institution/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUserIntstitution(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userIntUrl + "/institution/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
