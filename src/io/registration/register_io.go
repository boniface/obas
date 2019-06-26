package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/registration"
)

const registerUrl = api.BASE_URL + "/registration"

type Register domain.Register

func GetRegisters() ([]domain.Register, error) {
	entites := []domain.Register{}
	resp, _ := api.Rest().Get(registerUrl + "/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetRegister(id string) (domain.Register, error) {
	entity := domain.Register{}
	resp, _ := api.Rest().Get(registerUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateRegister(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(registerUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateRegister(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(registerUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteRegister(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(registerUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
