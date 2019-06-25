package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/application"
)

const applicationTypeUrl = api.BASE_URL + "/application"

type ApplicationType domain.ApplicationType

func GetApplicationTypes() ([]ApplicationType, error) {
	entites := []ApplicationType{}
	resp, _ := api.Rest().Get(applicationTypeUrl + "/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetApplication(id string) (ApplicationType, error) {
	entity := ApplicationType{}
	resp, _ := api.Rest().Get(applicationTypeUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateApplication(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationTyprUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func UpdateApplication(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationTyprUrl + "/update")
	if resp.IsError() {
		return true, nil
	}
	return true, nil
}
func DeleteApplication(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationTyprUrl + "delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
