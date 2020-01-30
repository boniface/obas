package demographics

import (
	"errors"
	"obas/api"
	domain "obas/domain/demographics"
)

const genderUrl = api.BASE_URL + "/demographics/gender/"

type Gender domain.Gender

func GetGenders() ([]Gender, error) {
	entites := []Gender{}
	//entites = append(entites, Gender{"1", "Male"})
	//entites = append(entites, Gender{"2", "Female"})
	resp, _ := api.Rest().Get(genderUrl + "all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetGender(id string) (Gender, error) {
	entity := Gender{}
	resp, _ := api.Rest().Get(genderUrl + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateGender(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(genderUrl + "create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateGender(entity interface{}, token string) (bool, error) {
	resp, _ := api.Rest().
		SetAuthToken(token).
		SetBody(entity).
		Post(genderUrl + "update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteGender(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(genderUrl + "delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
