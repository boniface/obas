package institutions

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/institutions"
)

const schoolUrl = api.BASE_URL + "/institutions"

type Schools domain.School

func GetSchools() ([]domain.School, error) {
	entites := []domain.School{}
	resp, _ := api.Rest().Get(schoolUrl + "/school/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetSchool(id string) (domain.School, error) {
	entity := domain.School{}
	resp, _ := api.Rest().Get(schoolUrl + "/school/get/" + id)
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
		Post(schoolUrl + "/school/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateSchool(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(schoolUrl + "/school/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func DeleteSchool(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(schoolUrl + "/school/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
