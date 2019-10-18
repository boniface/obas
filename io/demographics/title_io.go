package demographics

import (
	"errors"
	"obas/api"
	demographyDomain "obas/domain/demographics"
)

const titleUrl = api.BASE_URL + "/demographics"

type Titles demographyDomain.Title

func GetTitles() ([]Titles, error) {
	entites := []Titles{}
	resp, _ := api.Rest().Get(titleUrl + "/title/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetTitle(id string) (demographyDomain.Title, error) {
	entity := demographyDomain.Title{}
	resp, _ := api.Rest().Get(titleUrl + "/title/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateTitle(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(titleUrl + "/title/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateTitle(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(titleUrl + "/title/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteTitle(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(titleUrl + "/title/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
