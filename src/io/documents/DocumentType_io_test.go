package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDocumentTypes(t *testing.T) {
	value, err := GetDocument("")
	assert.Nil(t, err)
	assert.Equal(t, value, "entity", "Return entity")
}

func TestGetDocumentsTypes(t *testing.T) {
	value, err := GetDocument("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestCreateDocumentTypes(t *testing.T) {
	value, err := CreateDocument("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestUpdateDocumentTypes(t *testing.T) {
	value, err := UpdateDocument("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestDeleteDocumentTypes(t *testing.T) {
	value, err := DeleteDocument("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}
