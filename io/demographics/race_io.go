package demographics

import (
	"errors"
	"obas/api"
	demographyDomain "obas/domain/demographics"
)

const raceUrl = api.BASE_URL + "/demographics"

type Race demographyDomain.Race

func GetRaces() ([]Race, error) {
	entites := []Race{}
	//entites = append(entites, Race{"1", "Black"})
	//entites = append(entites, Race{"2", "White"})
	//entites = append(entites, Race{"3", "Colored"})
	resp, _ := api.Rest().Get(raceUrl + "/race/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetRace(id string) (demographyDomain.Race, error) {
	entity := demographyDomain.Race{}
	resp, _ := api.Rest().Get(raceUrl + "/race/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateRace(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(raceUrl + "/race/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateRace(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(raceUrl + "/race/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteRace(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(raceUrl + "/race/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
