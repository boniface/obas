package documents

import (
	"errors"
	"obas/api"
	domain "obas/domain/documents"
)

const documentTypeUrl = api.BASE_URL + "/documents"

type DocumentType domain.DocumentType

func GetDocumentTypes() ([]DocumentType, error) {
	entites := []DocumentType{}
	//d1 := DocumentType{"1", "Matric"}
	//d2 := DocumentType{"2", "Identification"}
	//entites = []DocumentType{d1, d2}
	resp, _ := api.Rest().Get(documentTypeUrl + "/type/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetDocumentType(id string) (DocumentType, error) {
	entity := DocumentType{}
	resp, _ := api.Rest().Get(documentTypeUrl + "/type/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateDocumentType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(documentUrl + "/type/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateDocumentType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(documentUrl + "/type/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteDocumentType(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(documentUrl + "/type/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
