package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/users"
)

const processingTypeSatusUrl = api.BASE_URL + "/users"

type ProcessingStatusType domain.ProcessingStatusType

func GetProcessingStatusTypes() ([]ProcessingStatusType, error) {
	entites := []ProcessingStatusType{}
	resp, _ := api.Rest().Get(processingTypeSatusUrl + "/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetProcessingStatusType(id string) (ProcessingStatusType, error) {
	entity := ProcessingStatusType{}
	resp, _ := api.Rest().Get(processingTypeSatusUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateProcessingStatusType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(processingTypeSatusUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateProcessingStatusType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(processingTypeSatusUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteProcessingStatustype(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(processingTypeSatusUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
