package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userapplicationstatus = api.BASE_URL + "/users/applicationstatus"

func CreateUserApplicationStatus(obj domain.UserApplicationStatus) (domain.UserApplicationStatus, error) {
	entity := domain.UserApplicationStatus{}
	resp, _ := api.Rest().SetBody(obj).Post(userapplicationURL + "/create")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err == nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserApplicationStatus(id string) (domain.UserApplicationStatus, error) {
	entity := domain.UserApplicationStatus{}
	resp, _ := api.Rest().Get(userapplicationURL + "/get" + id)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err == nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserApplicationStatues() ([]domain.UserApplicationStatus, error) {
	entity := []domain.UserApplicationStatus{}
	resp, _ := api.Rest().Get(userapplicationURL + "/all")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err == nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteApplicationStatus(obj domain.UserApplicationStatus) (domain.UserApplicationStatus, error) {
	entity := domain.UserApplicationStatus{}
	resp, _ := api.Rest().SetBody(obj).Post(userapplicationURL + "/delete")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err == nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateUserApplicationStatus(obj domain.UserApplicationStatus) (domain.UserApplicationStatus, error) {
	entity := domain.UserApplicationStatus{}
	resp, _ := api.Rest().SetBody(obj).Post(userapplicationURL + "/update")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err == nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
