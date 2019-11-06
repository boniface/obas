package storage

import (
	"errors"
	"obas/api"
	domain "obas/domain/storage"
)

const storageURL = api.BASE_URL + "/file/"

type FileData domain.FileData

func UploadFile(filepath string, token string) (FileData, error) {
	entity := FileData{}
	resp, _ := api.Rest().
		SetAuthToken(token).
		SetFile("upload", filepath).
		Post(storageURL + "upload")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
