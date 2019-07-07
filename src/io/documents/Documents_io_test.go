package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDocument(t *testing.T) {
	value, err := GetDocument("12132")
	assert.Nil(t, err)
	assert.Equal(t, value, "hgkjhkj", "this should be ....")
}

func TestGetDocuments(t *testing.T) {
	value, err := GetDocument("er")
	assert.NotNil(t, err)
	assert.Equal(t, value, "d")
}

func TestCreateDocument(t *testing.T) {
	value, err := CreateDocument("f")
	assert.NotNil(t, err)
	assert.Equal(t, value, "as")
}

func TestUpdateDocument(t *testing.T) {
	value, err := UpdateDocument("y")
	assert.NotNil(t, err)
	assert.Equal(t, value, "as")
}

func TestDeleteDocument(t *testing.T) {
	value, err := DeleteDocument("e")
	assert.NotNil(t, err)
	assert.Equal(t, value, "as")
}
