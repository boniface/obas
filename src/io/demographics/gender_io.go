package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/demographics"
)

const genderUrl = api.BASE_URL + "/demographics"

type Genders domain.Gender

func GetGenders() ([]Genders, error) {
	entites := []Genders{}
	resp, _ := api.Rest().Get(genderUrl + "/gender/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetGender(id string) (Genders, error) {
	entity := Genders{}
	resp, _ := api.Rest().Get(genderUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateGender(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(genderUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateGender(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(genderUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteGender(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(genderUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
