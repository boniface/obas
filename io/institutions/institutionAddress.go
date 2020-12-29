package institutions

import (
	"errors"
	"obas/api"
	domain "obas/domain/institutions"
)

const institutionAddressURl = api.BASE_URL + "/institutions/address/"

func CreateInstitutionAddress(obj domain.InstitutionAddress) (domain.InstitutionAddress, error) {
	entity := domain.InstitutionAddress{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionAddressURl + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadInstitutionAddress(institutionId, addressTypeId string) (domain.InstitutionAddress, error) {
	entity := domain.InstitutionAddress{}
	resp, _ := api.Rest().Get(institutionAddressURl + "get/" + institutionId + "/" + addressTypeId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetInstitutionAddresses(institutionId string) ([]domain.InstitutionAddress, error) {
	entity := []domain.InstitutionAddress{}
	resp, _ := api.Rest().Get(institutionAddressURl + "getaddresses/" + institutionId)
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
	resp, _ := api.Rest().SetBody(obj).Post(institutionAddressURl + "delete")
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
	resp, _ := api.Rest().SetBody(obj).Post(institutionAddressURl + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetAllInstitutionAddresses() ([]domain.InstitutionAddress, error) {
	entity := []domain.InstitutionAddress{}
	resp, _ := api.Rest().Get(institutionAddressURl + "all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
