package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/users"
)

const studentProfileUrl = api.BASE_URL + "/users"

type StudentProfiles domain.StudentProfile

func GetStudentProfiles() ([]domain.StudentProfile, error) {
	entites := []domain.StudentProfile{}
	resp, _ := api.Rest().Get(studentProfileUrl + "/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetStudentProfile(id string) (domain.StudentProfile, error) {
	entity := domain.StudentProfile{}
	resp, _ := api.Rest().Get(studentProfileUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateStudentProfiles(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentProfileUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateStudentProfiles(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentProfileUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteStudentProfiles(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentProfileUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
