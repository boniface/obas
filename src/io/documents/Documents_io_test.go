package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDocuments(t *testing.T) {

	value, err := GetDocuments()
	assert.Nil(t, err)
	//assert.Equal(t, value, "hgkjhkj", "this should be ....")
	assert.True(t, len(value) > 0)
}

func TestGetDocument(t *testing.T) {
	expected := ""
	value, err := GetDocument("")
	assert.Nil(t, err)
	assert.Equal(t, expected, value)
}

func TestCreateDocument(t *testing.T) {
	value, err := CreateDocument("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateDocument(t *testing.T) {
	value, err := UpdateDocument("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestDeleteDocument(t *testing.T) {
	value, err := DeleteDocument("")
	assert.Nil(t, err)
	assert.True(t, value)
}
