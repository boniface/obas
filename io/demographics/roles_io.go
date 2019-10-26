package demographics

import (
	"errors"
	"obas/api"
	domain "obas/domain/demographics"
)

const roleUrl = api.BASE_URL + "/demographics"

type Role domain.Role

func GetRoles() ([]Role, error) {
	entites := []Role{}
	resp, _ := api.Rest().Get(roleUrl + "/roles/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetRole(id string) (Role, error) {
	entity := Role{}
	resp, _ := api.Rest().Get(roleUrl + "/roles/get/" + id)
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
		Post(roleUrl + "/roles/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateRole(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(roleUrl + "/roles/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteRole(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(roleUrl + "/roles/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
