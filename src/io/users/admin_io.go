package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/users"
)

const adminUrl = api.BASE_URL + "/users"

type Admin domain.Admin

func GetAdmins() ([]Admin, error) {
	entites := []Admin{}
	resp, _ := api.Rest().Get(adminUrl + "/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetAdmin(id string) (domain.Admin, error) {
	entity := domain.Admin{}
	resp, _ := api.Rest().Get(adminUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateAdmin(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(adminUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateAdmin(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(adminUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteAdmin(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(adminUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
