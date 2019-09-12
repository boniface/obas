package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/users"
	"testing"
)

func TestGetUserRoles(t *testing.T) {
	result, err := GetUserRoles()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, len(result) > 0)
}

func TestGetUserRole(t *testing.T) {
	expected := "25"
	result, err := GetUserRole("30")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, expected, result.RoleId)

}

func TestCreateUserRole(t *testing.T) {
	role := domain.UserRole{"35", "ADMIN"}
	result, err := CreateUserRole(role)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)

}

func TestUpdateUserRole(t *testing.T) {
	role := domain.UserRole{"35", "48"}
	result, err := UpdateUserRole(role)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)
}

func TestDeleteUserRole(t *testing.T) {
	role := domain.UserRole{"34", "48"}
	result, err := DeleteUserRole(role)
	assert.Nil(t, err)
	assert.True(t, result)

}
