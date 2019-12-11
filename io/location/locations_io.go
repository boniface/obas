package location

import (
	"errors"
	"fmt"
	"obas/api"
	domain "obas/domain/location"
)

const locationUrl = api.BASE_URL + "/location"

func GetLocations() ([]domain.Location, error) {
	entites := []domain.Location{}
	resp, serverEr := api.Rest().Get(locationUrl + "/all")
	if resp.IsError() {
		fmt.Println(" Is request from Server Okay", serverEr)
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetLocation(id string) (domain.Location, error) {
	entity := domain.Location{}
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

func CreateLocation(entity domain.Location) (domain.Location, error) {
	location := domain.Location{}
	resp, _ := api.Rest().
		SetBody(entity).
		Post(locationUrl + "/create")
	if resp.IsError() {
		return location, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &location)
	if err != nil {
		return location, errors.New(resp.Status())
	}

	return location, nil
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
