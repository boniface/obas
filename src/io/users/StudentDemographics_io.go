package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/users"
)

const studentDemographicsUrl = api.BASE_URL + "/users"

type StudentDemographics domain.StudentDemographics

func GetStudentDemographics() ([]StudentDemographics, error) {
	entites := []StudentDemographics{}
	resp, _ := api.Rest().Get(studentDemographicsUrl + "/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetStudentDemographic(id string) (StudentDemographics, error) {
	entity := StudentDemographics{}
	resp, _ := api.Rest().Get(studentDemographicsUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateStudentDemographics(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentDemographicsUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateStudentDemographics(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentDemographicsUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteStudentDemographics(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentDemographicsUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
