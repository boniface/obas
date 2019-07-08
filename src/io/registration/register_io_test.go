package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisters(t *testing.T) {
	result, err := GetRegisters()
	assert.Nil(t, err)

	assert.True(t, len(result) > 0)

}

func TestRegister(t *testing.T) {
	expected := ""
	result, err := GetRegister("")
	assert.Nil(t, err)
	assert.Equal(t, expected, result.Email)

}

func TestCreateRegister(t *testing.T) {
	result, err := CreateRegister("")
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestUpdateRegister(t *testing.T) {
	result, err := UpdateRegister("")
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDeleteRegister(t *testing.T) {
	result, err := DeleteRegister("")
	assert.Nil(t, err)
	assert.True(t, result)

}
