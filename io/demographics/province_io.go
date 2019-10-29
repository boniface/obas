package demographics

import (
	"errors"
	"obas/api"
	domain "obas/domain/demographics"
)

const provinceURL = api.BASE_URL + "/demographics/province"

type Province domain.Province

func GetProvinces()([]Province, error) {
	entities := []Province{}
	//entities = append(entities, Province{"1", "Eastern Cape"})
	//entities = append(entities, Province{"2", "Western Cape"})
	resp, _ := api.Rest().Get(provinceURL + "/all")

	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}

func GetProvince(provinceCode string) (Province, error) {
	entity := Province{}
	resp, _ := api.Rest().Get(provinceURL + "/get/" + provinceCode)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}