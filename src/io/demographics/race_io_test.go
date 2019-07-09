package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRaces(t *testing.T) {
	value, err := GetRaces()
	assert.Nil(t, err)
	assert.True(t, len(value) > 0)
}

func TestGetRace(t *testing.T) {
	expected := ""
	value, err := GetRace("")
	assert.Nil(t, err)
	assert.Equal(t, value, expected)
}

func TestCreateRace(t *testing.T) {
	value, err := CreateRace("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateRace(t *testing.T) {
	value, err := UpdateRace("")
	assert.Nil(t, err)
	assert.NotNil(t, value)
}

func TestDeleteRace(t *testing.T) {
	value, err := DeleteRace("")
	assert.Nil(t, err)
	assert.NotNil(t, value)
}
