package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/users"
)

const studentDocumentsUrl = api.BASE_URL + "/users"

type StudentDocuments domain.StudentDocuments

func GetStudentDocuments() ([]domain.StudentDocuments, error) {
	entites := []domain.StudentDocuments{}
	resp, _ := api.Rest().Get(studentDocumentsUrl + "/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetStudentDocument(id string) (domain.StudentDocuments, error) {
	entity := domain.StudentDocuments{}
	resp, _ := api.Rest().Get(studentDocumentsUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateStudentDocuments(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentDocumentsUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateStudentDocuments(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentDocumentsUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteStudentDocuments(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentDocumentsUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
