package io

import (
	"encoding/json"
	"errors"
	"obas/src/api"
	domain "obas/src/domain/mail"
)

const smtpUrl = api.BASE_URL + "/smtp"

type SmtpConfig domain.SmtpConfig

func GetSmtpConfigs() ([]SmtpConfig, error) {
	entities := []SmtpConfig{}
	resp, _ := api.Rest().Get(smtpUrl + "/smtp/all")
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil

}

func GetSmtpConfig(id string) (SmtpConfig, error) {
	role := SmtpConfig{}
	resp, _ := api.Rest().Get(smtpUrl + "/smtp/get/" + id)
	if resp.IsError() {
		return role, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &role)
	if err != nil {
		return role, errors.New(resp.Status())
	}
	return role, nil

}

func CreateSmtpConfig(entity SmtpConfig) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(smtpUrl + "/smtp/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}
func UpdateSmtpConfig(entity SmtpConfig) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(smtpUrl + "/smtp/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}

func DeleteSmtpConfig(entity SmtpConfig) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(smtpUrl + "/smtp/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil

}
