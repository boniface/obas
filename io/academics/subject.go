package academics

import (
	"errors"
	"obas/api"
	domain "obas/domain/academics"
)

//please add url
const subjectURL = api.BASE_URL

func CreateSubject(obj domain.Subject) (domain.Subject, error) {
	entity := domain.Subject{}
	resp, _ := api.Rest().SetBody(obj).Post(subjectURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetSubject(id string) (domain.Subject, error) {
	entity := domain.Subject{}
	resp, _ := api.Rest().Get(subjectURL + "/get" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetSubjects(id string) ([]domain.Subject, error) {
	entity := []domain.Subject{}
	resp, _ := api.Rest().Get(subjectURL + "/all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteSubject(obj domain.Subject) (domain.Subject, error) {
	entity := domain.Subject{}
	resp, _ := api.Rest().SetBody(obj).Post(subjectURL + "/delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateSubject(obj domain.Subject) (domain.Subject, error) {
	entity := domain.Subject{}
	resp, _ := api.Rest().SetBody(obj).Post(subjectURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
