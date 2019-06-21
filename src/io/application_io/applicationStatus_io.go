package application_io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/application"
)

const applicationStatusUrl = api.BASE_URL + "/application"

type applicationStatus domain.ApplicationStatus

func getApplicationStatuses() ([]applicationStatus, error) {
	entites := []applicationStatus{}
	resp, _ := api.Rest().get(applicationStatusUrl + "/all")
	if resp.isError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func getApplicationStatus(id string) (applicationStatus, error) {
	entity := applicationStatus{}
	resp, _ := api.Rest().get(applicationStatusUrl + "/get/" + id)
	if resp.isError() {
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
		post(applicationStatusUrl + "/create")
	if resp.IsError() {
		return false, erroes.New(resp.Status())
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
func DeleteApplicationStatus(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationStatusUrl + "delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
