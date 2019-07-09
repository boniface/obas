package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDocuments(t *testing.T) {
	value, _ := GetDocuments()
	assert.Nil(t, value)
}

func TestGetDocument(t *testing.T) {
	value, _ := GetDocument("")
	assert.NotNil(t, value)
	assert.Equal(t, value, "Return entity")
}

func TestCreateDocument(t *testing.T) {
	value, _ := CreateDocument("")
	assert.NotNil(t, value)
	assert.Equal(t, value, "Return entity")
}

func TestUpdateDocument(t *testing.T) {
	value, _ := UpdateDocument("")
	assert.NotNil(t, value)
	assert.Equal(t, value, "Return entity")
}

func TestDeleteDocument(t *testing.T) {
	value, _ := DeleteDocument("")
	assert.NotNil(t, value)
	assert.Equal(t, value, "Return entity")
}
