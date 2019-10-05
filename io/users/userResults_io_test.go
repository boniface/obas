package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

func TestGetUserResults(t *testing.T) {
	result, err := GetUserResults()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, len(result) > 0)
}

func TestGetUserResult(t *testing.T) {
	expected := "PASS"
	result, err := GetUserResult("12")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, expected, result.Description)

}

func TestCreateUserResults(t *testing.T) {
	uResults := domain.UserResults{"12", "PASS"}
	result, err := CreateUserResults(uResults)
	assert.Nil(t, err)
	assert.True(t, result)

}

func TestUpdateUserResults(t *testing.T) {
	uResults := domain.UserResults{"12", "PASS"}
	result, err := UpdateUserResults(uResults)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)
}

func TestDeleteUserResults(t *testing.T) {
	uResults := domain.UserResults{"12", "PASS"}
	result, err := DeleteUserResults(uResults)
	assert.Nil(t, err)
	assert.True(t, result)

}
