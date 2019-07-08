package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRaces(t *testing.T) {
	value, _ := GetRaces()
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestGetRace(t *testing.T) {
	value, _ := GetRace("121")
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestCreateRace(t *testing.T) {
	value, _ := CreateRace(0)
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestUpdateRace(t *testing.T) {
	value, _ := UpdateRace(0)
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestDeleteRace(t *testing.T) {
	value, _ := DeleteRace(0)
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}
