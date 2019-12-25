package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userMatricSubjectURL = api.BASE_URL + "/institution/matric/subject/"

func CreateUserMatricSubject(obj domain.UserMatricSubject) (domain.UserMatricSubject, error) {
	entity := domain.UserMatricSubject{}
	resp, _ := api.Rest().SetBody(obj).Post(userMatricSubjectURL + "create")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserMatricSubjectForUser(userId string) ([]domain.UserMatricSubject, error) {
	entity := []domain.UserMatricSubject{}
	resp, _ := api.Rest().Get(userMatricSubjectURL + "allforuser/" + userId)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserMatricSubjects() ([]domain.UserMatricSubject, error) {
	entity := []domain.UserMatricSubject{}
	resp, _ := api.Rest().Get(userMatricSubjectURL + "all")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserMatricSubject(userId, subjectId string) ([]domain.UserMatricSubject, error) {
	entity := []domain.UserMatricSubject{}
	resp, _ := api.Rest().Get(userMatricSubjectURL + "get/" + userId + "/" + subjectId)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateUserMatricSubject(obj domain.UserMatricSubject) (domain.UserMatricSubject, error) {
	entity := domain.UserMatricSubject{}
	resp, _ := api.Rest().SetBody(obj).Post(userMatricSubjectURL + "update")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteUserMatricSubject(obj domain.UserMatricSubject) (domain.UserMatricSubject, error) {
	entity := domain.UserMatricSubject{}
	resp, _ := api.Rest().SetBody(obj).Post(userMatricSubjectURL + "delete")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
