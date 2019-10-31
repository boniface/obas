package storage

import (
	"errors"
	"mime/multipart"
	"obas/api"
	domain "obas/domain/storage"
)

const storageURL = api.BASE_URL + "/file/"

type FileData domain.FileData

func UploadFile(file multipart.File, token string) (FileData, error) {
	entity := FileData{}
	resp, _ := api.Rest().
		SetAuthToken(token).
		SetBody(entity).
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
