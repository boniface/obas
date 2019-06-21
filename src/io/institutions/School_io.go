package institutions

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/demographics"
)

const schoolUrl = api.BASE_URL + "/demographics"

type Schools domain.Schools

func GetSchools() ([]Schools, error) {
	entites := []Schools{}
	resp, _ := api.Rest().Get(schoolUrl + "/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetSchool(id string) (Schools, error) {
	entity := Schools{}
	resp, _ := api.Rest().Get(schoolUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateSchool(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(schoolUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateSchool(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(schoolUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func DeleteSchool(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(schoolUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
