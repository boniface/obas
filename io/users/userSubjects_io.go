package users

import (
	"errors"
	"fmt"
	"obas/api"
	domain "obas/domain/users"
)

const userSubUrl = api.BASE_URL + "/users"

type uSubjects domain.UserSubjects

func GetUserSubjects() ([]uSubjects, error) {
	entites := []uSubjects{}
	resp, _ := api.Rest().Get(userSubUrl + "/subjects/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUserSubject(id string) (uSubjects, error) {
	entity := uSubjects{}
	resp, _ := api.Rest().Get(userSubUrl + "/subjects/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserSubject(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userSubUrl + "/subjects/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUserSubject(entity interface{}) (bool, error) {
	resp, serverEr := api.Rest().
		SetBody(entity).
		Post(userSubUrl + "/subjects/update")
	if resp.IsError() {
		fmt.Println(" Is request from Server Okay", serverEr)
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUserSubject(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userSubUrl + "/subjects/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
