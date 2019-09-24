package io

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userContUrl = api.BASE_URL + "/users"

type UserContacts domain.UserContacts

func GetUserContacts() ([]UserContacts, error) {
	entites := []UserContacts{}
	resp, _ := api.Rest().Get(userContUrl + "/contacts/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUserContact(id string) (UserContacts, error) {
	entity := UserContacts{}
	resp, _ := api.Rest().Get(userContUrl + "/contacts/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserContact(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userContUrl + "/contacts/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUserContact(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userContUrl + "/contacts/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUserContact(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userContUrl + "/contacts/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
