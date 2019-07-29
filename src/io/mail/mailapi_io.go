package io

import (
	"encoding/json"
	"errors"
	"obas/src/api"
	domain "obas/src/domain/mail"
)

const mailapiUrl = api.BASE_URL + "/mail"

type MailApi domain.MailApi

func GetMailApis() ([]MailApi, error) {
	entities := []MailApi{}
	resp, _ := api.Rest().Get(mailapiUrl + "/api/all")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil

}

func GetMailApi(id string) (domain.MailApi, error) {
	entity := domain.MailApi{}
	resp, _ := api.Rest().Get(mailapiUrl + "/api/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}

func CreateMailApi(entity domain.MailApi) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(mailapiUrl + "/api/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}
func UpdateMailApi(entity domain.MailApi) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(mailapiUrl + "/api/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}

func DeleteMailApi(entity domain.MailApi) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(mailapiUrl + "/api/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}
