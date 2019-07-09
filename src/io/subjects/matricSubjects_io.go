package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/subjects"
)

const matricSubjectUrl = api.BASE_URL + "/subjects"

type MatricSubjects domain.MatricSubjects

func GetMatricSubjects() ([]MatricSubjects, error) {
	entites := []MatricSubjects{}
	resp, _ := api.Rest().Get(matricSubjectUrl + "/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetMatricSubject(id string) (MatricSubjects, error) {
	entity := MatricSubjects{}
	resp, _ := api.Rest().Get(matricSubjectUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateMatricSubject(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(matricSubjectUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateMatricSubject(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(matricSubjectUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteMatricSubject(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(matricSubjectUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
