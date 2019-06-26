package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/subjects"
)

const subjectUrl = api.BASE_URL + "/subjects"

type Subjects domain.Subjects

func GetSubjects() ([]domain.Subjects, error) {
	entites := []domain.Subjects{}
	resp, _ := api.Rest().Get(subjectUrl + "/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetSubject(id string) (domain.Subjects, error) {
	entity := domain.Subjects{}
	resp, _ := api.Rest().Get(subjectUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateSubject(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(subjectUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateSubject(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(subjectUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteSubject(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(subjectUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
