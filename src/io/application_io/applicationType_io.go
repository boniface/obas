package application_io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/application"
)

const applicationTyprUrl = api.BASE_URL + "/application"

type application domain.ApplicationType

func getApplicationTypes() ([]application, error) {
	entites := []application{}
	resp, _ := api.Rest().get(applicationTyprUrl + "/all")
	if resp.isError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func getApplication(id string) (application, error) {
	entity := application{}
	resp, _ := api.Rest().get(applicationTyprUrl + "/get/" + id)
	if resp.isError() {
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
		post(applicationTyprUrl + "/create")
	if resp.IsError() {
		return false, erroes.New(resp.Status())
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
