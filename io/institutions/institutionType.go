package institutions

import (
	"errors"
	"obas/api"
	domain "obas/domain/institutions"
)

const institutionType = api.BASE_URL + "/institution/institution"

func CreateInstitutionType(obj domain.InstitutionTypes) (domain.InstitutionTypes, error) {
	entity := domain.InstitutionTypes{}
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
func GetInstitutionType(id string) (domain.InstitutionTypes, error) {
	entity := domain.InstitutionTypes{}
	resp, _ := api.Rest().Get(institutionAddressURl + "/get")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetInstitutionTypes() ([]domain.InstitutionTypes, error) {
	entity := []domain.InstitutionTypes{}
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
func DeleteInstitutionType(obj domain.InstitutionTypes) (domain.InstitutionTypes, error) {
	entity := domain.InstitutionTypes{}
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
func UpdateInstitutionType(obj domain.InstitutionTypes) (domain.InstitutionTypes, error) {
	entity := domain.InstitutionTypes{}
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
