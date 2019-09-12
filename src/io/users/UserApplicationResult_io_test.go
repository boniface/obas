package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/users"
	"testing"
)

var appResult = domain.UserApplicationResult{"22", "DECISION COLLECTION"}

func TestGetUserApplicationResults(t *testing.T) {
	result, err := GetUserApplicationResults()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, len(result) > 0)
}

func TestGetUserApplicationResult(t *testing.T) {
	expected := appResult
	result, err := GetUserApplicationResult(appResult.UserApplicationResultId)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

}

func TestCreateUserApplicationResult(t *testing.T) {
	result, err := CreateUserApplicationResult(appResult)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)

}

func TestUpdateUserApplicationResult(t *testing.T) {
	var expected = "PENDING"
	var updated = domain.UserApplicationResult{"22", "PENDING"}
	result, err := UpdateUserApplicationResult(updated)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)
	value, err := GetUserApplicationResult(appResult.UserApplicationResultId)
	assert.Equal(t, expected, value.Description)
}

func TestDeleteUserApplicationResult(t *testing.T) {
	appResult := domain.UserApplicationResult{"29", "SUCCESSFUL"}
	result, err := DeleteUserApplicationResult(appResult)
	assert.Nil(t, err)
	assert.True(t, result)

}
