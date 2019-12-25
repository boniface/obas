package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userTertiaryInstitutionURL = api.BASE_URL + "/institution/tertiary/"

func GetUserTertiaryInstitutions() ([]domain.UserTertiaryInstitution, error) {
	entity := []domain.UserTertiaryInstitution{}
	resp, _ := api.Rest().Get(userTertiaryInstitutionURL + "all")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetAllForUser(userId string) (domain.UserTertiaryInstitution, error) {
	entity := domain.UserTertiaryInstitution{}
	resp, _ := api.Rest().Get(userTertiaryInstitutionURL + "allforuser/" + userId)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserTertiaryInstitutionForApp(userId, application string) (domain.UserTertiaryInstitution, error) {
	entity := domain.UserTertiaryInstitution{}
	resp, _ := api.Rest().Get(userTertiaryInstitutionURL + "getforapplication/" + userId + "/" + application)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CreateUserTertiaryInstitution(obj domain.UserTertiaryInstitution) (domain.UserTertiaryInstitution, error) {
	entity := domain.UserTertiaryInstitution{}
	resp, _ := api.Rest().SetBody(obj).Post(userTertiaryInstitutionURL + "create/")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateUserTertiaryInstitution(obj domain.UserTertiaryInstitution) (domain.UserTertiaryInstitution, error) {
	entity := domain.UserTertiaryInstitution{}
	resp, _ := api.Rest().SetBody(obj).Post(userTertiaryInstitutionURL + "update/")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteUserTertiaryInstitution(obj domain.UserTertiaryInstitution) (domain.UserTertiaryInstitution, error) {
	entity := domain.UserTertiaryInstitution{}
	resp, _ := api.Rest().SetBody(obj).Post(userTertiaryInstitutionURL + "detele/")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
