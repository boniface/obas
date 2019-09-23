package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/users"
)

const userDocUrl = api.BASE_URL + "/users"

type UserDocuments domain.UserDocuments

func GetUserDocuments() ([]UserDocuments, error) {
	entites := []UserDocuments{}
	resp, _ := api.Rest().Get(userDocUrl + "/documents/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUserDocument(id string) (UserDocuments, error) {
	entity := UserDocuments{}
	resp, _ := api.Rest().Get(userDocUrl + "/documents/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserDocument(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userDocUrl + "/documents/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUserDocument(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userDocUrl + "/documents/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUserDocument(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userDocUrl + "/documents/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
