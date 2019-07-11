package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLocations(t *testing.T) {
	value, err := GetLocations()
	assert.Nil(t, err)
	assert.True(t, len(value) > 0)
}

func TestGetLocation(t *testing.T) {
	expected := ""
	value, err := GetLocation("")
	assert.Nil(t, err)
	assert.Equal(t, value, expected)
}

func TestCreateSchool(t *testing.T) {
	value, err := CreateLocation("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateDocument(t *testing.T) {
	value, err := UpdateLocation("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestDeleteDocument(t *testing.T) {
	value, err := DeleteLocation("")
	assert.Nil(t, err)
	assert.True(t, value)
}
