package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUsers(t *testing.T) {
	result, err := GetUsers()
	assert.Nil(t, err)

	assert.True(t, len(result) > 0)
}

func TestGetUser(t *testing.T) {
	expected := ""
	result, err := GetUser("")
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

}

func TestCreateUser(t *testing.T) {
	result, err := CreateUser("")
	assert.Nil(t, err)
	assert.True(t, result)

}

func TestUpdateUser(t *testing.T) {
	result, err := UpdateUser("")
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDeleteUser(t *testing.T) {
	result, err := DeleteUser("")
	assert.Nil(t, err)
	assert.True(t, result)

}
