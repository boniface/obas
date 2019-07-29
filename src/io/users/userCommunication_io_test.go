package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/users"
	"testing"
)

func TestGetUserCommunications(t *testing.T) {
	value, err := GetUserCommunications()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetUserCommunication(t *testing.T) {
	expected := "COURIER"
	value, err := GetUserCommunication("42")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.Description, expected)
}

func TestCreateUserCommunication(t *testing.T) {
	uCom := domain.UserCommunication{"63", "EMAILS"}
	value, err := CreateUserCommunication(uCom)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestUpdateUserCommunication(t *testing.T) {
	uCom := domain.UserCommunication{"63", "EMAILS"}
	value, err := UpdateUserCommunication(uCom)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteUserCommunication(t *testing.T) {
	uCom := domain.UserCommunication{"65", "POSTAL"}
	value, err := DeleteUserCommunication(uCom)
	assert.Nil(t, err)
	assert.True(t, value)
}
