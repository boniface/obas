package institutions

import (
	"errors"
	"obas/api"
	domain "obas/domain/institutions"
)

const institutionLocationURL = api.BASE_URL + "/institutions"

func CreateInstitutionLocation(obj domain.InstitutionLocation) (domain.InstitutionLocation, error) {
	entity := domain.InstitutionLocation{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionLocationURL + "/location/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteInstitutionLocation(obj domain.InstitutionLocation) (domain.InstitutionLocation, error) {
	entity := domain.InstitutionLocation{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionLocationURL + "/location/delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadInstitutionLocation(id string) (domain.InstitutionLocation, error) {
	entity := domain.InstitutionLocation{}
	resp, _ := api.Rest().Get(institutionLocationURL + "/location/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadInstitutionLocations() ([]domain.InstitutionLocation, error) {
	entity := []domain.InstitutionLocation{}
	resp, _ := api.Rest().Get(institutionLocationURL + "/location/all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateInstitutionLocation(obj domain.InstitutionLocation) (domain.InstitutionLocation, error) {
	entity := domain.InstitutionLocation{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionLocationURL + "/location/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
