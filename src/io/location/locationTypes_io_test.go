package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLocationTypes(t *testing.T) {
	value, err := GetLocationTypes()
	assert.Nil(t, err)
	assert.True(t, len(value) > 0)
}

func TestGetLocationType(t *testing.T) {
	expected := ""
	value, err := GetLocationType("")
	assert.Nil(t, err)
	assert.Equal(t, value, expected)
}

func TestCreateLocationType(t *testing.T) {
	value, err := CreateLocationType("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateLocationType(t *testing.T) {
	value, err := UpdateLocationType("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestDeleteLocationType(t *testing.T) {
	value, err := DeleteLocationType("")
	assert.Nil(t, err)
	assert.True(t, value)
}
