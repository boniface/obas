package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/registration"
)



const EmailUrl = api.BASE_URL + "/registration"

type Email domain.Roles

func GetEmail() ([]Email, error) {
	entites := []Email{}
	resp, _ := api.Rest().Get(emailUrl + "/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetEmail(id string) (Email, error) {
	entity := Roles{}
	resp, _ := api.Rest().Get(emailUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateEmail(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(emailUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateEmail(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(emailUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteEmail(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(emailUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
