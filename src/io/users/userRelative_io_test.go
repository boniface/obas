package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/users"
	"testing"
)

func TestGetUserRelatives(t *testing.T) {
	result, err := GetUserRelatives()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, len(result) > 0)
}

func TestGetUserRelative(t *testing.T) {
	expected := "JEAN"
	result, err := GetUserRelative("7896541230")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, expected, result.Name)

}

func TestCreateUserRelative(t *testing.T) {
	relative := domain.UserRelative{"36", "JOSH", "073321456850", "BROTHERS", "acl@gogo.com"}
	result, err := CreateUserRelative(relative)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)

}

func TestUpdateUserRelative(t *testing.T) {
	relative := domain.UserRelative{"36", "JOSH", "073321456850", "BROTHERS", "acl@gogo.com"}
	result, err := UpdateUserRelative(relative)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)
}

func TestDeleteUserRelative(t *testing.T) {
	relative := domain.UserRelative{"m@gt.com", "JEAN", "PAUL", "MATUTO", "m@h.com"}
	result, err := DeleteUserRelative(relative)
	assert.Nil(t, err)
	assert.True(t, result)

}
