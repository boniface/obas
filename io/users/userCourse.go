package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const usercourseURL = api.BASE_URL + "/users"

func CreateUserCourse(obj domain.UserCourse) (domain.UserCourse, error) {
	entity := domain.UserCourse{}
	resp, _ := api.Rest().SetBody(obj).Post(usercourseURL + "/course/create")
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
	resp, _ := api.Rest().Get(usercourseURL + "/course/all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserCourse(userId, institutionId, courseId string) (domain.UserCourse, error) {
	entity := domain.UserCourse{}
	resp, _ := api.Rest().Get(usercourseURL + "/course/get/" + userId + "/" + institutionId + "/" + courseId)
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
	resp, _ := api.Rest().SetBody(obj).Post(usercourseURL + "/course/delete")
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
	resp, _ := api.Rest().SetBody(obj).Post(usercourseURL + "/course/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
