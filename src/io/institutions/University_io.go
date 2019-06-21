package institutions

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/demographics"
)

const universityUrl = api.BASE_URL + "/demographics"

type Universitys domain.Universitys

func GetUniversitys() ([]Universitys, error) {
	entites := []Universitys{}
	resp, _ := api.Rest().Get(universityUrl + "/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUniversity(id string) (Universitys, error) {
	entity := Universitys{}
	resp, _ := api.Rest().Get(universityUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUniversity(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(universityUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUniversity(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(universityUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUniversity(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(universityUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
