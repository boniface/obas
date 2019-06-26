package institutions

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/institutions"
)

const universityUrl = api.BASE_URL + "/institutions"

type Universitys domain.University

func GetUniversitys() ([]domain.University, error) {
	entites := []domain.University{}
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

func GetUniversity(id string) (domain.University, error) {
	entity := domain.University{}
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
