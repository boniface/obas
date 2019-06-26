package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/application"
)

const applicationResultUrl = api.BASE_URL + "/application"

type ApplicationResult domain.ApplicationResult

func GetApplicationResultes() ([]domain.ApplicationResult, error) {
	entites := []domain.ApplicationResult{}
	resp, _ := api.Rest().Get(applicationResultUrl + "/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetApplicationResult(id string) (domain.ApplicationResult, error) {
	entity := domain.ApplicationResult{}
	resp, _ := api.Rest().Get(applicationResultUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateApplicationResult(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationResultUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func UpdateApplicationResult(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationResultUrl + "/update")
	if resp.IsError() {
		return true, nil
	}
	return true, nil
}
func DeleteApplicationResult(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationResultUrl + "delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
