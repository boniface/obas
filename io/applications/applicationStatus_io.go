package applications

import (
	"errors"
	"obas/api"
	domain "obas/domain/application"
)

const applicationStatusUrl = api.BASE_URL + "/application"

type ApplicationStatus domain.ApplicationStatus

func GetApplicationStatuses() ([]ApplicationStatus, error) {
	entites := []ApplicationStatus{}
	resp, _ := api.Rest().Get(applicationStatusUrl + "/status/all")
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
	resp, _ := api.Rest().Get(applicationStatusUrl + "/status/get/" + id)
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
		Post(applicationStatusUrl + "/status/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func UpdateApplicationStatus(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationStatusUrl + "/status/update")
	if resp.IsError() {
		return true, nil
	}
	return true, nil
}
func DeleteApplicationStatus(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationStatusUrl + "/status/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
