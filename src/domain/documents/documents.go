package domain

import "time"

type Documents struct {
	Email          string    `json:"email"`
	DocumentsId    string    `json:"documentsId"`
	DocumentTypeId string    `json:"documentTypeId"`
	Description    string    `json:"description"`
	Url            string    `json:"url"`
	Mime           string    `json:"mime"`
	Date           time.Time `json:"date"`
	Permission     string    `json:"permission"`
}

type DocumentType struct {
	DocumentTypeId   string `json:"documentTypeId"`
	DocumentTypeName string `json:" documentTypeName"`
}
