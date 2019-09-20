package io

import (
	"encoding/json"
	"errors"
	"obas/src/api"
	domain "obas/src/domain/security"
)

const apiUrl = api.BASE_URL + "/security"

type ApiKeys domain.ApiKeys

func getApiKey(id string) (ApiKeys, error) {
	apikey := ApiKeys{}
	resp, _ := api.Rest().Get(apiUrl + "/get/" + id)
	if resp.IsError() {
		return apikey, errors.New(resp.Status())
	}
	err := json.Unmarshal(resp.Body(), &apikey)
	if err != nil {
		return apikey, errors.New(resp.Status())
	}
	return apikey, nil

}
