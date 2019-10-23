package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userAddressUrl = api.BASE_URL + "/users"

type UserAddress domain.UserAddress

func GetUserAddresses() ([]UserAddress, error) {
	entities := []UserAddress{}
	//addT1 := UserAddress{"caniksea@yahoo.co.nz", "123", "81 Loop Street", "8001"}
	//addT2 := UserAddress{"caniksea@yahoo.co.nz", "246", "P.O.Box 3245 Brackenfell", "6792"}
	//
	//allAdd := []UserAddress{addT1, addT2}
	//
	//entities = allAdd
	resp, _ := api.Rest().Get(userAddressUrl + "/address/all")

	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}

func GetUserAddress(userId string, addressTypeId string) (UserAddress, error) {
	entity := UserAddress{}
	//if addressTypeId == "123" {
	//	entity = UserAddress{userId, addressTypeId, "81 Loop Street", "8001"}
	//} else if addressTypeId == "246" {
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

func UpdateUserAddress(entity UserAddress, token string) (bool, error) {
	resp, _ := api.Rest().
		SetAuthToken(token).
		SetBody(entity).
		Post(userAddressUrl + "/address/update")
	if resp.IsError() {
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
