package location

import (
	"errors"
	"obas/api"
	domain "obas/domain/location"
)

const locationUrl = api.BASE_URL + "/location"

type Location domain.Location

func GetLocations() ([]domain.Location, error) {
	entites := []domain.Location{}
	resp, _ := api.Rest().Get(locationUrl + "/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetLocation(id string) (Location, error) {
	entity := Location{}
	resp, _ := api.Rest().Get(locationUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateLocation(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(locationUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateLocation(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(locationUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteLocation(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(locationUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
