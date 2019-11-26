package academics

import (
	"errors"
	"obas/api"
	domain "obas/domain/academics"
)

//please add the URL part
const courseSubjectURL = api.BASE_URL

func CreateCourseSubject(obj domain.CourseSubject) (domain.CourseSubject, error) {
	entity := domain.CourseSubject{}

	resp, _ := api.Rest().SetBody(obj).Post(courseSubjectURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetCourseSubject(id string) (domain.CourseSubject, error) {
	entity := domain.CourseSubject{}
	resp, _ := api.Rest().Get(courseSubjectURL + "/get" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetCourseSubjects() ([]domain.CourseSubject, error) {
	entity := []domain.CourseSubject{}
	resp, _ := api.Rest().Get(courseSubjectURL + "/all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteCourseSubject(obj domain.CourseSubject) (domain.CourseSubject, error) {
	entity := domain.CourseSubject{}
	resp, _ := api.Rest().SetBody(obj).Post(courseSubjectURL + "/delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
