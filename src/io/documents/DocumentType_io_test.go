package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDocumentTypes(t *testing.T) {
	value, err := GetDocumentTypes()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetDocumentsType(t *testing.T) {
	expected := ""
	value, err := GetDocument("")
	assert.Nil(t, err)
	assert.Equal(t, value, expected)
}

func TestCreateDocumentTypes(t *testing.T) {
	value, err := CreateDocument("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateDocumentTypes(t *testing.T) {
	value, err := UpdateDocument("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestDeleteDocumentTypes(t *testing.T) {
	value, err := DeleteDocument("")
	assert.Nil(t, err)
	assert.True(t, value)
}
