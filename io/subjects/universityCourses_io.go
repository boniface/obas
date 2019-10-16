package subjects

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/subjects"
)

const univCourseUrl = api.BASE_URL + "/subjects"

type UniversityCourses domain.UniversityCourses

func GetUniversityCourses() ([]UniversityCourses, error) {
	entites := []UniversityCourses{}
	resp, _ := api.Rest().Get(univCourseUrl + "/university/all")

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
	resp, _ := api.Rest().Get(univCourseUrl + "/university/get/" + id)
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
		Post(univCourseUrl + "/university/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUniversityCourses(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(univCourseUrl + "/university/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUniversityCourses(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(univCourseUrl + "/university/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
