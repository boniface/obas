package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/demographics"
)

const roleUrl = api.BASE_URL + "/demographics"

type Roles domain.Roles

func GetRoles() ([]Roles, error) {
	entites := []Roles{}
	resp, _ := api.Rest().Get(roleUrl + "/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetRole(id string) (Roles, error) {
	entity := Roles{}
	resp, _ := api.Rest().Get(roleUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateRole(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(roleUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateRole(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(roleUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteRole(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(roleUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
