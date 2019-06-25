package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/users"
)

const studentApplicationSatusUrl = api.BASE_URL + "/users"

type StudentApplicationStatus domain.StudentApplicationStatus

func GetStudentApplicationStatuses() ([]StudentApplicationStatus, error) {
	entites := []StudentApplicationStatus{}
	resp, _ := api.Rest().Get(studentApplicationSatusUrl + "/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetStudentApplicationStatus(id string) (StudentApplicationStatus, error) {
	entity := StudentApplicationStatus{}
	resp, _ := api.Rest().Get(studentApplicationSatusUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateStudentApplicationStatus(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentApplicationSatusUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateStudentApplicationStatus(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentApplicationSatusUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteStudentApplicationStatus(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentApplicationSatusUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
