package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/subjects"
)

const universityCoursesUrl = api.BASE_URL + "/subjects"

type universityCourses domain.UniversityCourses

func GetUniversityCourses() ([]UniversityCourses, error) {
	entites := []UniversityCourses{}
	resp, _ := api.Rest().Get(universityCoursesUrl + "/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUniversityCourse(id string) (UniversityCourses, error) {
	entity := UniversityCourses{}
	resp, _ := api.Rest().Get(universityCoursesUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUniversityCourses(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(universityCoursesUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUniversityCourses(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(universityCoursesUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUniversityCourses(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(universityCoursesUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
