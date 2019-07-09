package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLocationTypes(t *testing.T) {
	value, err := GetLocationTypes()
	assert.Nil(t, err)
	assert.Equal(t, value, "entity", "Return entity")
}

func TestGetLocationType(t *testing.T) {
	value, err := GetLocationType("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestCreateLocationType(t *testing.T) {
	value, err := CreateLocationType("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestUpdateLocationType(t *testing.T) {
	value, err := UpdateLocationType("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestDeleteLocationType(t *testing.T) {
	value, err := DeleteLocationType("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}
