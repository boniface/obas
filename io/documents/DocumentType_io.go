package documents

import (
	"errors"
	"obas/api"
	domain "obas/domain/documents"
)

const documentTypeUrl = api.BASE_URL + "/documents/type/"

func GetDocumentTypes() ([]domain.DocumentType, error) {
	entites := []domain.DocumentType{}
	//d1 := domain.DocumentType{"1", "Matric"}
	//d2 := domain.DocumentType{"2", "Identification"}
	//entites = []domain.DocumentType{d1, d2}
	resp, _ := api.Rest().Get(documentTypeUrl + "all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	println(resp.Body())
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetDocumentType(documentTypeId string) (domain.DocumentType, error) {
	entity := domain.DocumentType{}
	//if documentTypeId == "1" {
	//	entity = domain.DocumentType{documentTypeId, "Matric"}
	//} else if documentTypeId == "2" {
	//	entity = domain.DocumentType{documentTypeId, "ID"}
	//}
	resp, _ := api.Rest().Get(documentTypeUrl + "get/" + documentTypeId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateDocumentType(entity domain.DocumentType, token string) (bool, error) {
	resp, _ := api.Rest().SetAuthToken(token).SetBody(entity).
		Post(documentTypeUrl + "create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateDocumentType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(documentTypeUrl + "update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteDocumentType(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(documentTypeUrl + "delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
