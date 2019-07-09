package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDocumentTypes(t *testing.T) {
	value, _ := GetDocument("")
	assert.NotNil(t, value)
}

func TestGetDocumentsTypes(t *testing.T) {
	value, _ := GetDocument("")
	assert.NotNil(t, value)
	//assert.Equal(t, value, "Return entity")
}

func TestCreateDocumentTypes(t *testing.T) {
	value, _ := CreateDocument("")
	assert.NotNil(t, value)
	assert.Equal(t, value, "Return entity")
}

func TestUpdateDocumentTypes(t *testing.T) {
	value, _ := UpdateDocument("")
	assert.NotNil(t, value)
	assert.Equal(t, value, "Return entity")
}

func TestDeleteDocumentTypes(t *testing.T) {
	value, _ := DeleteDocument("")
	assert.NotNil(t, value)
	assert.Equal(t, value, "Return entity")
}
