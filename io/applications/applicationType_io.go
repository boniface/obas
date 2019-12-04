package applications

import (
	"errors"
	"fmt"
	"obas/api"
	domain "obas/domain/application"
)

const applicationTypeUrl = api.BASE_URL + "/application"

type ApplicationType domain.ApplicationType

func GetApplicationTypes() ([]ApplicationType, error) {
	entites := []ApplicationType{}
	resp, _ := api.Rest().Get(applicationTypeUrl + "/type/all")
	if resp.IsError() {
		fmt.Println(" Is request from Server Okay")
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		fmt.Println("Did Json Coversion Take Place Okay", err)
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetApplicationType(id string) (domain.ApplicationType, error) {
	entity := domain.ApplicationType{}
	resp, _ := api.Rest().Get(applicationTypeUrl + "/type/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateApplicationType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationTypeUrl + "/type/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func UpdateApplicationType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationTypeUrl + "/type/update")
	if resp.IsError() {
		return true, nil
	}
	return true, nil
}
func DeleteApplicationType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationTypeUrl + "/type/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
