package io

import (
	"errors"
	"fmt"
	"obas/src/api"
	domain "obas/src/domain/address"
)

const addressTypeUrl = api.BASE_URL + "/address"

type AddressType domain.AddressType

func GetAddressTypes() ([]AddressType, error) {
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

func GetAddressType(id string) (domain.AddressType, error) {
	entity := domain.AddressType{}
	resp, serverEr := api.Rest().Get(addressTypeUrl + "/get/" + id)
	if resp.IsError() {
		fmt.Println(" Is request from Server Okay", serverEr)
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		fmt.Println("Did Jason Coversion Take Place Okay", err)
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateAddressType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(addressTypeUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateAddressType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(addressTypeUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteAddressType(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(addressTypeUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
