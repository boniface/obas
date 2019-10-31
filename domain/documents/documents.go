package domain

import "time"

type Document struct {
	DocumentId    string `json:"documentId"`
	DocumentTypeId string `json:"documentTypeId"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	Mime           string `json:"mime"`
	Date           time.Time `json:"date"`
	Permission     string `json:"permission"`
	DocumentStatus string `json:"documentStatus"`
}

type DocumentType struct {
	DocumentTypeId   string `json:"documentTypeId"`
	DocumentTypeName string `json:" documentTypeName"`
}
