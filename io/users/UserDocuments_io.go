package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userDocUrl = api.BASE_URL + "/users"

type UserDocument domain.UserDocument

func GetUserDocuments(userId string) ([]UserDocument, error) {
	entities := []UserDocument{}
	resp, _ := api.Rest().Get(userDocUrl + "/documents/get/" + userId)

	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}

func GetUserDocument(id string) (UserDocument, error) {
	entity := UserDocument{}
	resp, _ := api.Rest().Get(userDocUrl + "/documents/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserDocument(entity UserDocument, token string) (bool, error) {
	resp, _ := api.Rest().
		SetAuthToken(token).
		SetBody(entity).
		Post(userDocUrl + "/documents/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUserDocument(entity UserDocument, token string) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userDocUrl + "/documents/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUserDocument(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userDocUrl + "/documents/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
