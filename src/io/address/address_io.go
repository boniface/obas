package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/address"
)

const addressUrl = api.BASE_URL + "/address"

type Address domain.Address

func GetAddresses() ([]Address, error) {
	entites := []Address{}
	resp, _ := api.Rest().Get(addressUrl + "/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetAddress(id string) (Address, error) {
	entity := Address{}
	resp, _ := api.Rest().Get(addressUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateAddress(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(addressUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateAddress(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(addressUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteAddress(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(addressUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
