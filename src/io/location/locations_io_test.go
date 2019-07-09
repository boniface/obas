package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLocations(t *testing.T) {
	value, err := GetLocations()
	assert.Nil(t, err)
	assert.Equal(t, value, "entity", "Return entity")
}

func TestGetLocation(t *testing.T) {
	value, err := GetLocation("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestCreateSchool(t *testing.T) {
	value, err := CreateLocation("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestUpdateDocument(t *testing.T) {
	value, err := UpdateLocation("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestDeleteDocument(t *testing.T) {
	value, err := DeleteLocation("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}
