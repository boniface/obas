package demographics

import (
	"errors"
	"obas/api"
	domain "obas/domain/demographics"
)

const townURL = api.BASE_URL + "/demographics/town"
const districtTownURL = api.BASE_URL + "/demographics/districttown"

type Town domain.Town
type DistictTown domain.DistrictTown

func GetTowns() ([]Town, error) {
	entities := []Town{}
	//entities = append(entities, Town{"1", "Cape Town"})
	//entities = append(entities, Town{"2", "Stellenbosch"})
	//entities = append(entities, Town{"3", "Breede Valley"})
	//entities = append(entities, Town{"4", "Cederberg"})
	//entities = append(entities, Town{"5", "Saldanabay"})
	resp, _ := api.Rest().Get(townURL + "/all")

	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}

func GetTown(townCode string) (Town, error) {
	entity := Town{}
	//if townCode == "1" {
	//	entity = Town{townCode, "Cape Town"}
	//} else if townCode == "2" {
	//	entity = Town{townCode, "Stellenbosch"}
	//} else if townCode == "3" {
	//	entity = Town{townCode, "Breede Valley"}
	//} else if townCode == "4" {
	//	entity = Town{townCode, "Cederberg"}
	//} else if townCode == "5" {
	//	entity = Town{townCode, "Saldanabay"}
	//}
	resp, _ := api.Rest().Get(townURL + "/get/" + townCode)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetTownsInDistrict(districtCode string) ([]DistictTown, error) {
	entities := []DistictTown{}
	//if districtCode == "1" {
	//	entities = append(entities, DistictTown{districtCode, "1"})
	//} else if districtCode == "2" {
	//	entities = append(entities, DistictTown{districtCode, "2"})
	//	entities = append(entities, DistictTown{districtCode, "3"})
	//} else if districtCode == "3" {
	//	entities = append(entities, DistictTown{districtCode, "4"})
	//	entities = append(entities, DistictTown{districtCode, "5"})
	//}
	resp, _ := api.Rest().Get(districtTownURL + "/gettowns/" + districtCode)
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, err
	}
	return entities, nil
}

func GetDistrictForTown(townCode string) (DistictTown, error) {
	entity := DistictTown{}
	//entity = DistictTown{"1", townCode}
	resp, _ := api.Rest().Get(districtTownURL + "/gettowndistrict/" + townCode)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
