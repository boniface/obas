package demographics

import (
	"errors"
	"obas/api"
	domain "obas/domain/demographics"
)

const provincedistrictURL = api.BASE_URL

func CreateProvinceDistrict(obj domain.ProvinceDistrict) (domain.ProvinceDistrict, error) {
	entity := domain.ProvinceDistrict{}
	resp, _ := api.Rest().SetBody(obj).Post(provincedistrictURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetProvinceDistricts() ([]domain.ProvinceDistrict, error) {
	entity := []domain.ProvinceDistrict{}
	resp, _ := api.Rest().Get(provincedistrictURL + "/all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetProvinceDistrict(id string) (domain.ProvinceDistrict, error) {
	entity := domain.ProvinceDistrict{}
	resp, _ := api.Rest().Get(provincedistrictURL + "/get" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteProvinceDistrict(obj domain.ProvinceDistrict) (domain.ProvinceDistrict, error) {
	entity := domain.ProvinceDistrict{}
	resp, _ := api.Rest().SetBody(obj).Post(provincedistrictURL + "/delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateProvinceDistrict(obj domain.ProvinceDistrict) (domain.ProvinceDistrict, error) {
	entity := domain.ProvinceDistrict{}
	resp, _ := api.Rest().SetBody(obj).Post(provincedistrictURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
