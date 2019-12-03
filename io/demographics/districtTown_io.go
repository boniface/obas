package demographics

import (
	"errors"
	"obas/api"
	domain "obas/domain/demographics"
)

//please complete the URL
const districttownURL = api.BASE_URL

func CreateDistrictTown(obj domain.DistrictTown) (domain.DistrictTown, error) {
	entity := domain.DistrictTown{}
	resp, _ := api.Rest().SetBody(obj).Post(districttownURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetDistrictTown(id string) (domain.DistrictTown, error) {
	entity := domain.DistrictTown{}
	resp, _ := api.Rest().Get(districttownURL + "id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetDistrictTowns() (domain.DistrictTown, error) {
	entity := domain.DistrictTown{}
	resp, _ := api.Rest().Get(districttownURL + "all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteDistrictTown(obj domain.DistrictTown) (domain.DistrictTown, error) {
	entity := domain.DistrictTown{}
	resp, _ := api.Rest().SetBody(obj).Post(districttownURL + "delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateDistrictTown(obj domain.DistrictTown) (domain.DistrictTown, error) {
	entity := domain.DistrictTown{}
	resp, _ := api.Rest().SetBody(obj).Post(districttownURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
