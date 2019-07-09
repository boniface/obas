package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/application"
)

const applicationStatusUrl = api.BASE_URL + "/application"

type ApplicationStatus domain.ApplicationStatus

func GetApplicationStatuses() ([]ApplicationStatus, error) {
	entites := []ApplicationStatus{}
	resp, _ := api.Rest().Get(applicationStatusUrl + "/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetApplicationStatus(id string) (ApplicationStatus, error) {
	entity := ApplicationStatus{}
	resp, _ := api.Rest().Get(applicationStatusUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateApplicationStatus(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationStatusUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func UpdateApplicationStatus(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationStatusUrl + "/update")
	if resp.IsError() {
		return true, nil
	}
	return true, nil
}
func DeleteApplicationStatus(id string) (ApplicationStatus, error) {
	entity := ApplicationStatus{}
	resp, _ := api.Rest().Get(applicationStatusUrl + "/delete/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
