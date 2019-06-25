package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/address"
)

const addressTypeUrl = api.BASE_URL + "/address"

type AddressType domain.AddressType

func GetAddresses() ([]AddressType, error) {
	entites := []AddressType{}
	resp, _ := api.Rest().Get(addressTypeUrl + "/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetAddress(id string) (AddressType, error) {
	entity := AddressType{}
	resp, _ := api.Rest().Get(addressTypeUrl + "/get/" + id)
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
		Post(addressTypeUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateAddress(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(addressTypeUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteAddress(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(addressTypeUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
