package institutions

import (
	"errors"
	"obas/api"
	domain "obas/domain/institutions"
)

const institutionURL = api.BASE_URL + "/institutions"

func CreateInstitution(obj domain.Institution) (domain.Institution, error) {
	entity := domain.Institution{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteInstitution(obj domain.Institution) (domain.Institution, error) {
	entity := domain.Institution{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionURL + "/delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetInstitution(id string) (domain.Institution, error) {
	entity := domain.Institution{}
	resp, _ := api.Rest().Get(institutionURL + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetInstitutions() ([]domain.Institution, error) {
	entity := []domain.Institution{}
	entity = append(entity, domain.Institution{"1", "1", "CPUT"})
	entity = append(entity, domain.Institution{"2", "1", "UWC"})
	entity = append(entity, domain.Institution{"3", "1", "Stellenbosch"})
	//resp, _ := api.Rest().Get(institutionURL + "/all")
	//if resp.IsError() {
	//	return entity, errors.New(resp.Status())
	//}
	//err := api.JSON.Unmarshal(resp.Body(), &entity)
	//if err != nil {
	//	return entity, errors.New(resp.Status())
	//}
	return entity, nil
}

func UpdateInstitution(obj domain.Institution) (domain.Institution, error) {
	entity := domain.Institution{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
