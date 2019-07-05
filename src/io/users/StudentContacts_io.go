package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/users"
)

const studentContactsUrl = api.BASE_URL + "/users"

type StudentContacts domain.StudentContacts

func GetStudentContacts() ([]StudentContacts, error) {
	entites := []StudentContacts{}
	resp, _ := api.Rest().Get(studentContactsUrl + "/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetStudentContact(id string) (StudentContacts, error) {
	entity := StudentContacts{}
	resp, _ := api.Rest().Get(studentContactsUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateStudentContact(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentContactsUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateStudentContact(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentContactsUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteStudentContact(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(studentContactsUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
