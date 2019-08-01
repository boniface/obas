package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/users"
	"testing"
)

func TestGetUserApplicationResults(t *testing.T) {
	result, err := GetUserApplicationResults()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, len(result) > 0)
}

func TestGetUserApplicationResult(t *testing.T) {
	expected := "SUCCESS"
	result, err := GetUserApplicationResult("24")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, expected, result.Description)

}

func TestCreateUserApplicationResult(t *testing.T) {
	appResult := domain.UserApplicationResult{"25", "SUCCESSFUL"}
	result, err := CreateUserApplicationResult(appResult)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)

}

func TestUpdateUserApplicationResult(t *testing.T) {
	appResult := domain.UserApplicationResult{"25", "Pending"}
	result, err := UpdateUserApplicationResult(appResult)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)
}

func TestDeleteUserApplicationResult(t *testing.T) {
	appResult := domain.UserApplicationResult{"25", "SUCCESSFUL"}
	result, err := DeleteUserApplicationResult(appResult)
	assert.Nil(t, err)
	assert.True(t, result)

}
