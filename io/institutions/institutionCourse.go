package institutions

import (
	"errors"
	"obas/api"
	domain "obas/domain/institutions"
)

const institutioncourseURL = api.BASE_URL

func GetInstitutionCourse(id string) (domain.InstitutionCourse, error) {
	entity := domain.InstitutionCourse{}
	resp, _ := api.Rest().Get(institutionAddressURl + "/get" + id)
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
	resp, _ := api.Rest().Get(institutionAddressURl + "/all")
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
	resp, _ := api.Rest().SetBody(obj).Post(institutionAddressURl + "/delete")
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
	resp, _ := api.Rest().SetBody(obj).Post(institutionAddressURl + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpadteInstitutionCourse(obj domain.InstitutionCourse) (domain.InstitutionCourse, error) {
	entity := domain.InstitutionCourse{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionAddressURl + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
