package users

import (
	"errors"
	"fmt"
	"obas/api"
	domain "obas/domain/users"
)

const userAddressUrl = api.BASE_URL + "/users"

type UserAddress domain.UserAddress

func GetUserAddresses() ([]UserAddress, error) {
	entites := []UserAddress{}
	resp, _ := api.Rest().Get(userAddressUrl + "/address/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUserAddress(userId string, addressTypeId string) (UserAddress, error) {
	entity := UserAddress{}
	//if addressTypeId == "1" {
	//	entity = UserAddress{userId, addressTypeId, "81 Loop Street", "8001"}
	//} else if addressTypeId == "2" {
	//	entity = UserAddress{userId, addressTypeId, "P.0.Box 3278 Brackenfell", "6792"}
	//}
	resp, _ := api.Rest().Get(userAddressUrl + "/address/get/" + userId + "/" + addressTypeId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserAddress(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userAddressUrl + "/address/create")
	if resp.IsError() {

		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateUserAddress(entity interface{}) (bool, error) {
	resp, serverEr := api.Rest().
		SetBody(entity).
		Post(userAddressUrl + "/address/update")
	if resp.IsError() {
		fmt.Println(" Is request from Server Okay", serverEr)
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUserAddress(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userAddressUrl + "/address/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
