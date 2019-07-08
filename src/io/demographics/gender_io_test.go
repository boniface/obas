package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetGenders(t *testing.T) {
	value, _ := GetGenders()
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestGetGender(t *testing.T) {
	value, _ := GetGender("121")
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestCreateGender(t *testing.T) {
	value, _ := CreateGender(0)
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestUpdateGender(t *testing.T) {
	value, _ := UpdateGender(0)
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestDeleteGender(t *testing.T) {
	value, _ := DeleteGender("121")
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}
