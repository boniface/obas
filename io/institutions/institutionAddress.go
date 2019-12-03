package institutions

import (
	"errors"
	"obas/api"
	domain "obas/domain/institutions"
)

const institutionAddressURl = api.BASE_URL + "/institutions"

func CreateInstitutionAddress(obj domain.InstitutionAddress) (domain.InstitutionAddress, error) {
	entity := domain.InstitutionAddress{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionAddressURl + "/address/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadInstitutionAddress(institutionId, addressId string) (domain.InstitutionAddress, error) {
	entity := domain.InstitutionAddress{}
	resp, _ := api.Rest().Get(institutionAddressURl + "/address/get/" + institutionId + "/" + addressId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

//need to know what this method is returnning
func GetInstitutionAddress(institutionId string) (domain.InstitutionAddress, error) {
	entity := domain.InstitutionAddress{}
	resp, _ := api.Rest().Get(institutionAddressURl + "/address/getAddresses/" + institutionId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetInstitutionAddresses() ([]domain.InstitutionAddress, error) {
	entity := []domain.InstitutionAddress{}
	resp, _ := api.Rest().Get(institutionAddressURl + "/address/all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteInstitutionAddress(obj domain.InstitutionAddress) (domain.InstitutionAddress, error) {
	entity := domain.InstitutionAddress{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionAddressURl + "/address/delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateInstitutionAddress(obj domain.InstitutionAddress) (domain.InstitutionAddress, error) {
	entity := domain.InstitutionAddress{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionAddressURl + "/address/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
