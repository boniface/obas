package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/users"
)

const studentResultUrl = api.BASE_URL + "/users"

type StudentResults domain.StudentResults

func GetStudentResults() ([]domain.StudentResults, error) {
	entites := []domain.StudentResults{}
	resp, _ := api.Rest().Get(studentResultUrl + "/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetStudentResult(id string) (domain.StudentResults, error) {
	entity := domain.StudentResults{}
	resp, _ := api.Rest().Get(studentResultUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateStudentResults(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentResultUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateStudentResults(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentResultUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteStudentResults(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentResultUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
