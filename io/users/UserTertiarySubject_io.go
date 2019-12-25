package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userTertiarySubjectURL = api.BASE_URL + "/institution/tertiary/subject/"

func CreateUserTertiarySubject(obj domain.UserTertiarySubject) (domain.UserTertiarySubject, error) {
	entity := domain.UserTertiarySubject{}

	resp, _ := api.Rest().SetBody(obj).Post(userTertiarySubjectURL + "create")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserTertiarySubject(userId string) ([]domain.UserTertiarySubject, error) {
	entity := []domain.UserTertiarySubject{}

	resp, _ := api.Rest().Get(userTertiarySubjectURL + "allforuser/" + userId)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserTertiarySubjects() ([]domain.UserTertiarySubject, error) {
	entity := []domain.UserTertiarySubject{}

	resp, _ := api.Rest().Get(userTertiarySubjectURL + "all/")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserTertiarySubjectallforapp(userId, applicationId string) (domain.UserTertiarySubject, error) {
	entity := domain.UserTertiarySubject{}
	resp, _ := api.Rest().Get(userTertiarySubjectURL + "allforapplication/" + userId + "/" + applicationId)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserTertiarySubjectGetForApp(userId, applicationId, subjectId string) (domain.UserTertiarySubject, error) {
	entity := domain.UserTertiarySubject{}
	resp, _ := api.Rest().Get(userTertiarySubjectURL + "getforapplication/" + userId + "/" + applicationId + "/" + subjectId)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteUserApplTertiarySubject(userId, applicationId string) (domain.UserTertiarySubject, error) {
	entity := domain.UserTertiarySubject{}
	resp, _ := api.Rest().Get(userTertiarySubjectURL + "deleteforapplication/" + userId + "/" + applicationId)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateUserTertiarySubject(obj domain.UserTertiarySubject) (domain.UserTertiarySubject, error) {
	entity := domain.UserTertiarySubject{}

	resp, _ := api.Rest().SetBody(obj).Post(userTertiarySubjectURL + "update")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetDeleteUserTertiarySubject(obj domain.UserTertiarySubject) (domain.UserTertiarySubject, error) {
	entity := domain.UserTertiarySubject{}
	resp, _ := api.Rest().SetBody(obj).Post(userTertiarySubjectURL + "delete")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
