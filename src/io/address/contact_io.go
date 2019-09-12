package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/address"
)

const contactTypeUrl = api.BASE_URL + "/address"

type ContactType domain.ContactType

func GetContactTypes() ([]ContactType, error) {
	entites := []ContactType{}
	resp, _ := api.Rest().Get(contactTypeUrl + "/contact/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetContactType(id string) (domain.ContactType, error) {
	entity := domain.ContactType{}
	resp, _ := api.Rest().Get(contactTypeUrl + "/contact/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateContactType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(contactTypeUrl + "/contact/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateContactType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(contactTypeUrl + "/contact/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteContactType(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(contactTypeUrl + "/contact/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
