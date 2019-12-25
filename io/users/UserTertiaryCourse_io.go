package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userTertiaryCourseURL = api.BASE_URL + "/institution/tertiary/course/"

func CreateUserTertiaryCourse(obj domain.UserTertiaryCourse) (domain.UserTertiaryCourse, error) {
	entity := domain.UserTertiaryCourse{}
	resp, _ := api.Rest().SetBody(obj).Post(userTertiaryCourseURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserTertiaryCourse(userId string) ([]domain.UserTertiaryCourse, error) {
	entity := []domain.UserTertiaryCourse{}
	resp, _ := api.Rest().Get(userTertiaryCourseURL + "allforuser/" + userId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserTertiaryCourses() ([]domain.UserTertiaryCourse, error) {
	entity := []domain.UserTertiaryCourse{}
	resp, _ := api.Rest().Get(userTertiaryCourseURL + "all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserTertiaryCourseForApp(userId, applicationId string) (domain.UserTertiaryCourse, error) {
	entity := domain.UserTertiaryCourse{}
	resp, _ := api.Rest().Get(userTertiaryCourseURL + "getforapplication/" + userId + "/" + applicationId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateUserTertiaryCourse(obj domain.UserTertiaryCourse) (domain.UserTertiaryCourse, error) {
	entity := domain.UserTertiaryCourse{}
	resp, _ := api.Rest().SetBody(obj).Post(userTertiaryCourseURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteUserTertiaryCourse(obj domain.UserTertiaryCourse) (domain.UserTertiaryCourse, error) {
	entity := domain.UserTertiaryCourse{}
	resp, _ := api.Rest().SetBody(obj).Post(userTertiaryCourseURL + "delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
