package academics

import (
	"errors"
	"obas/api"
	domain "obas/domain/academics"
)

const academicsURL = api.BASE_URL + "/course/"

type Course domain.Course

func SaveCourse(entity Course) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(academicsURL + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func GetCourse(id string) (Course, error) {
	entity := Course{}
	resp, _ := api.Rest().Get(academicsURL + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateCourse(course Course, token string) (bool, error) {
	resp, _ := api.Rest().
		SetAuthToken(token).
		SetBody(course).
		Post(academicsURL + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func DeleteCourse(course Course) (bool, error) {
	resp, _ := api.Rest().
		SetBody(course).
		Post(academicsURL + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func GetAllCourses() ([]Course, error) {
	//entites := []Course{}
	//resp, _ := api.Rest().Get(academicsURL + "/all")
	//
	//if resp.IsError() {
	//	return entites, errors.New(resp.Status())
	//}
	//err := api.JSON.Unmarshal(resp.Body(), &entites)
	//if err != nil {
	//	return entites, errors.New(resp.Status())
	//}
	entities := []Course{}
	entities = append(entities, Course{"1", "IT", ""})
	entities = append(entities, Course{"2", "Busines", ""})
	entities = append(entities, Course{"3", "Education", ""})
	return entities, nil
}
