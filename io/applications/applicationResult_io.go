package applications

import (
	"errors"
	"obas/api"
	domain "obas/domain/application"
)

const applicationResultUrl = api.BASE_URL + "/bursary"

type ApplicationResult domain.ApplicationResult

func GetApplicationResults() ([]ApplicationResult, error) {
	entites := []ApplicationResult{}
	resp, _ := api.Rest().Get(applicationResultUrl + "/result/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetApplicationResult(id string) (ApplicationResult, error) {
	entity := ApplicationResult{}
	resp, _ := api.Rest().Get(applicationResultUrl + "/result/get/" + id)
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
		Post(applicationResultUrl + "/result/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func UpdateApplicationResult(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationResultUrl + "/result/update")
	if resp.IsError() {
		return true, nil
	}
	return true, nil
}
func DeleteApplicationResult(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationResultUrl + "/result/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
