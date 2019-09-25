package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/documents"
	"testing"
)

func TestGetDocumentTypes(t *testing.T) {
	value, err := GetDocumentTypes()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetDocumentsType(t *testing.T) {
	expected := "COURSES"
	value, err := GetDocument("2")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value, expected)
}

func TestCreateDocumentTypes(t *testing.T) {
	docType := domain.DocumentType{"2", "COURSES"}
	value, err := CreateDocument(docType)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateDocumentTypes(t *testing.T) {
	docType := domain.DocumentType{"2", "COURSES"}
	value, err := UpdateDocument(docType)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteDocumentTypes(t *testing.T) {
	docType := domain.DocumentType{"2", "COURSES"}
	value, err := DeleteDocumentType(docType)
	assert.Nil(t, err)
	assert.True(t, value)
}
