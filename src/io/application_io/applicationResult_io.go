package application_io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/application"
)

const applicationResultUrl = api.BASE_URL + "/application"

type applicationResult domain.ApplicationResult

func getApplicationResultes() ([]applicationResult, error) {
	entites := []applicationResult{}
	resp, _ := api.Rest().get(applicationResultUrl + "/all")
	if resp.isError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func getApplicationResult(id string) (applicationResult, error) {
	entity := applicationResult{}
	resp, _ := api.Rest().get(applicationResultUrl + "/get/" + id)
	if resp.isError() {
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
		post(applicationResultUrl + "/create")
	if resp.IsError() {
		return false, erroes.New(resp.Status())
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
