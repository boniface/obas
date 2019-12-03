package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userAppResult = api.BASE_URL + "/users"

type UserApplicationResult domain.UserApplicationResult

func GetUserApplicationResults() ([]UserApplicationResult, error) {
	entites := []UserApplicationResult{}
	resp, _ := api.Rest().Get(userAppResult + "/bursary/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUserApplicationResult(id string) (UserApplicationResult, error) {
	entity := UserApplicationResult{}
	resp, _ := api.Rest().Get(userAppResult + "/bursary/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserApplicationResult(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userAppResult + "/bursary/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUserApplicationResult(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userAppResult + "/bursary/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUserApplicationResult(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userAppResult + "/bursary/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
