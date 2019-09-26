package io

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const uDemographicsUrl = api.BASE_URL + "/users"

type uDemographics domain.UserDemographics

func GetUserDemographics() ([]uDemographics, error) {
	entites := []uDemographics{}
	resp, _ := api.Rest().Get(uDemographicsUrl + "/demographics/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUserDemographic(id string) (uDemographics, error) {
	entity := uDemographics{}
	resp, _ := api.Rest().Get(uDemographicsUrl + "/demographics/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserDemographics(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(uDemographicsUrl + "/demographics/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUserDemographics(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(uDemographicsUrl + "/demographics/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUserDemographics(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(uDemographicsUrl + "/demographics/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
