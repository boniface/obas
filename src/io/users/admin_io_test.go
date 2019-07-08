package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAdmins(t *testing.T) {
	result, err := GetAdmins()
	assert.Nil(t, err)

	assert.True(t, len(result) > 0)
}

func TestGetAdmin(t *testing.T) {
	expected := ""
	result, err := GetAdmin("")
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

}

func TestCreateAdmin(t *testing.T) {
	result, err := CreateAdmin("")
	assert.Nil(t, err)
	assert.True(t, result)

}

func TestUpdateAdmin(t *testing.T) {
	result, err := UpdateAdmin("")
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDeleteAdmin(t *testing.T) {
	result, err := DeleteAdmin("")
	assert.Nil(t, err)
	assert.True(t, result)

}
