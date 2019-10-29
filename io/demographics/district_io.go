package demographics

import (
	"errors"
	"obas/api"
	domain "obas/domain/demographics"
)

const districtURL = api.BASE_URL + "/demographics/district"
const provinceDistrictURL = api.BASE_URL + "/demographics/provincedistrict"

type District domain.District
type ProvinceDistrict domain.ProvinceDistrict

func GetDistricts() ([]District, error) {
	entities := []District{}
	//entities = append(entities, District{"1", "City of Cape Town"})
	//entities = append(entities, District{"2", "Cape Winelands"})
	//entities = append(entities, District{"3", "West Coast"})
	resp, _ := api.Rest().Get(districtURL + "/all")

	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}

func GetDistrict(districtCode string) (District, error) {
	entity := District{}
	//if districtCode == "1" {
	//	entity = District{"1", "City of Cape Town"}
	//} else if districtCode == "2" {
	//	entity = District{"2", "Cape Winelands"}
	//} else if districtCode == "3" {
	//	entity = District{"3", "West Coast"}
	//}
	resp, _ := api.Rest().Get(districtURL + "/get/" + districtCode)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetDistrictsInProvince(provinceCode string) ([]ProvinceDistrict, error) {
	entities := []ProvinceDistrict{}
	//if provinceCode == "2" {
	//	entities = append(entities, ProvinceDistrict{provinceCode, "1"})
	//	entities = append(entities, ProvinceDistrict{provinceCode, "2"})
	//	entities = append(entities, ProvinceDistrict{provinceCode, "3"})
	//}
	resp, _ := api.Rest().Get(provinceDistrictURL + "/get/" + provinceCode)
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, err
	}
	return entities, nil
}

func GetProvinceForDistrict(districtCode string) (ProvinceDistrict, error) {
	entity := ProvinceDistrict{}
	//entity = ProvinceDistrict{"2", districtCode}
	resp, _ := api.Rest().Get(provinceDistrictURL + "/get/province/" + districtCode)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
