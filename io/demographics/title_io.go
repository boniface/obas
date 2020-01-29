package demographics

import (
	"errors"
	"obas/api"
	demographyDomain "obas/domain/demographics"
)

const titleUrl = api.BASE_URL + "/demographics/title/"

type Title demographyDomain.Title

func GetTitles() ([]Title, error) {
	entites := []Title{}
	//entites = append(entites, Title{"1", "Mr"})
	//entites = append(entites, Title{"2", "Miss"})
	//entites = append(entites, Title{"3", "Mrs"})
	//entites = append(entites, Title{"4", "Dr"})
	resp, _ := api.Rest().Get(titleUrl + "all")

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
	resp, _ := api.Rest().Get(titleUrl + "get/" + id)
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
		Post(titleUrl + "create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateTitle(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(titleUrl + "update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteTitle(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(titleUrl + "delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
