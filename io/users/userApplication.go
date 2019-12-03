package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

//please add the url
const userapplicationURL = api.BASE_URL

func CreateUserApplication(obj domain.UserApplication) (domain.UserApplication, error) {
	entity := domain.UserApplication{}
	resp, _ := api.Rest().SetBody(obj).Post(userapplicationURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserApplication(id string) (domain.UserApplication, error) {
	entity := domain.UserApplication{}
	resp, _ := api.Rest().Get(userapplicationURL + "/get")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserApplications(id string) ([]domain.UserApplication, error) {
	entity := []domain.UserApplication{}
	resp, _ := api.Rest().Get(userapplicationURL + "/all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateUserApplication(obj domain.UserApplication) (domain.UserApplication, error) {
	entity := domain.UserApplication{}
	resp, _ := api.Rest().SetBody(obj).Post(userapplicationURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteUserApplication(obj domain.UserApplication) (domain.UserApplication, error) {
	entity := domain.UserApplication{}
	resp, _ := api.Rest().SetBody(obj).Post(userapplicationURL + "/delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
