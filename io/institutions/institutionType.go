package institutions

import (
	"errors"
	"obas/api"
	domain "obas/domain/institutions"
)

const institutionType = api.BASE_URL + "/institutions"

func CreateInstitutionType(obj domain.InstitutionTypes) (domain.InstitutionTypes, error) {
	entity := domain.InstitutionTypes{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionAddressURl + "/type/create")
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
	resp, _ := api.Rest().Get(institutionAddressURl + "/type/get/" + id)
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
	resp, _ := api.Rest().Get(institutionAddressURl + "/type/all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteInstitutionType(obj domain.InstitutionTypes) (bool, error) {
	//entity := domain.InstitutionTypes{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionAddressURl + "/type/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	//err := api.JSON.Unmarshal(resp.Body(), &entity)
	//if err != nil {
	//	return entity, errors.New(resp.Status())
	//}
	return true, nil
}
func UpdateInstitutionType(obj domain.InstitutionTypes) (domain.InstitutionTypes, error) {
	entity := domain.InstitutionTypes{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionAddressURl + "/type/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
