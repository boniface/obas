package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userContUrl = api.BASE_URL + "/users"

type UserContact domain.UserContacts

func GetUserContacts() ([]UserContact, error) {
	entites := []UserContact{}
	resp, _ := api.Rest().Get(userContUrl + "/contacts/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUserContact(userId string, contactTypeId string) (UserContact, error) {
	entity := UserContact{}
	//if contactTypeId == "1" {
	//	entity = UserContact{userId, contactTypeId, "0145235521"}
	//} else if contactTypeId == "2" {
	//	entity = UserContact{userId, contactTypeId, "alternative@gmail.com"}
	//}
	resp, _ := api.Rest().Get(userContUrl + "/contacts/get/" + userId + "/" + contactTypeId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserContact(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userContUrl + "/contacts/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUserContact(entity UserContact, token string) (bool, error) {
	resp, _ := api.Rest().
		SetAuthToken(token).
		SetBody(entity).
		Post(userContUrl + "/contacts/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUserContact(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userContUrl + "/contacts/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
