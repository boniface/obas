package institutions

import (
	"errors"
	"obas/api"
	domain "obas/domain/institutions"
)

const institutioncourseURL = api.BASE_URL + "/institution_course"

func ReadInstitutionCourse(institutionId, courseId string) (domain.InstitutionCourse, error) {
	entity := domain.InstitutionCourse{}
	resp, _ := api.Rest().Get(institutioncourseURL + "/course/get/" + institutionId + "/" + courseId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetInstitutionCourses() ([]domain.InstitutionCourse, error) {
	entity := []domain.InstitutionCourse{}
	resp, _ := api.Rest().Get(institutioncourseURL + "/course/all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteInstitutionCourse(obj domain.InstitutionCourse) (domain.InstitutionCourse, error) {
	entity := domain.InstitutionCourse{}
	resp, _ := api.Rest().SetBody(obj).Post(institutioncourseURL + "/course/delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CreateInstitutionCourse(obj domain.InstitutionCourse) (domain.InstitutionCourse, error) {
	entity := domain.InstitutionCourse{}
	resp, _ := api.Rest().SetBody(obj).Post(institutioncourseURL + "/course/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetInstitutionCourse(institutionId string) (domain.InstitutionCourse, error) {
	entity := domain.InstitutionCourse{}
	resp, _ := api.Rest().Get(institutioncourseURL + "/course/getcourses/" + institutionId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
