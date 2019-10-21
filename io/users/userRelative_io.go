package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userRelUrl = api.BASE_URL + "/users"

type UserRelative domain.UserRelative

func GetUserRelatives() ([]UserRelative, error) {
	entites := []UserRelative{}
	resp, _ := api.Rest().Get(userRelUrl + "/relative/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUserRelative(id string) (UserRelative, error) {
	entity := UserRelative{}
	//entity = UserRelative{"caniksea@yahoo.co.nz", "Isaac Anikwue", "0983828432", "", "Father"}
	resp, _ := api.Rest().Get(userRelUrl + "/relative/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserRelative(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userRelUrl + "/relative/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUserRelative(entity UserRelative, token string) (bool, error) {
	resp, _ := api.Rest().
		SetAuthToken(token).
		SetBody(entity).
		Post(userRelUrl + "/relative/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUserRelative(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userRelUrl + "/relative/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
