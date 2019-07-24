package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDocuments(t *testing.T) {
	value, err := GetDocuments()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetDocument(t *testing.T) {
	expected := ""
	value, err := GetDocument("")
	assert.Nil(t, err)
	assert.Equal(t, value, expected)
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
