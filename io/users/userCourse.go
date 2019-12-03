package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const usercourseURL = api.BASE_URL

func CreateUserCourse(obj domain.UserCourse) (domain.UserCourse, error) {
	entity := domain.UserCourse{}
	resp, _ := api.Rest().SetBody(obj).Post(usercourseURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserCourses() ([]domain.UserCourse, error) {
	entity := []domain.UserCourse{}
	resp, _ := api.Rest().Get(usercourseURL + "/all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserCourse() (domain.UserCourse, error) {
	entity := domain.UserCourse{}
	resp, _ := api.Rest().Get(usercourseURL + "/get")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteUserCourse(obj domain.UserCourse) (domain.UserCourse, error) {
	entity := domain.UserCourse{}
	resp, _ := api.Rest().SetBody(obj).Post(usercourseURL + "/delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateUserCourse(obj domain.UserCourse) (domain.UserCourse, error) {
	entity := domain.UserCourse{}
	resp, _ := api.Rest().SetBody(obj).Post(usercourseURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
