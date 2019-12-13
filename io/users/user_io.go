package users

import (
	"errors"
	"obas/api"
	userDomain "obas/domain/users"
	"time"
)

const usersUrl = api.BASE_URL + "/users"

type User userDomain.User

func GetUsers() ([]User, error) {
	entites := []User{}
	resp, _ := api.Rest().Get(usersUrl + "/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUser(id string) (User, error) {
	entity := User{}
	entity = User{id, "4829830090930", "Arinze", "", "Anikwue", time.Now()}
	//resp, _ := api.Rest().Get(usersUrl + "/get/" + id)
	//if resp.IsError() {
	//	return entity, errors.New(resp.Status())
	//}
	//err := api.JSON.Unmarshal(resp.Body(), &entity)
	//if err != nil {
	//	return entity, errors.New(resp.Status())
	//}
	return entity, nil
}

func CreateUser(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(usersUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUser(entity User, token string) (bool, error) {
	resp, _ := api.Rest().
		SetAuthToken(token).
		SetBody(entity).
		Post(usersUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUser(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(usersUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
