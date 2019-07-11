package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetGenders(t *testing.T) {
	value, err := GetGenders()
	assert.Nil(t, err)
	assert.NotNil(t, len(value) > 0)
}

func TestGetGender(t *testing.T) {
	expected := ""
	value, err := GetGender("")
	assert.Nil(t, err)
	assert.Equal(t, value, expected)
}

func TestCreateGender(t *testing.T) {
	value, err := CreateGender("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateGender(t *testing.T) {
	value, err := UpdateGender("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestDeleteGender(t *testing.T) {
	value, err := DeleteGender("")
	assert.Nil(t, err)
	assert.True(t, value)
}
