package address_io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/demographics"
)

const contactTypeUrl = api.BASE_URL + "/address"

type contactType domain.Roles

func GetRoles() ([]contactType, error) {
	entites := []contactType{}
	resp, _ := api.Rest().Get(contactTypeUrl + "/all")
	if resp.isError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetRole(id string) (contactType, error) {
	entity := contactType{}
	resp, _ := api.Rest().Get(contactTypeUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateRole(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(contactTypeUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateRole(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(contactTypeUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteRole(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(contactTypeUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
