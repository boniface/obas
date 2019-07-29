package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/users"
	"testing"
)

func TestGetUserSubjects(t *testing.T) {
	result, err := GetUserSubjects()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, len(result) > 0)
}

func TestGetUserSubject(t *testing.T) {
	expected := "BUSINESS"
	result, err := GetUserSubject("35")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, expected, result.Name)

}

func TestCreateUserSubject(t *testing.T) {
	Subject := domain.UserSubjects{"35", "BUSINESS", "THEORY", "FIRST"}
	result, err := CreateUserSubject(Subject)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)

}

func TestUpdateUserSubject(t *testing.T) {
	Subject := domain.UserSubjects{"35", "BUSINESS", "THEORY", "FIRST"}
	result, err := UpdateUserSubject(Subject)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)
}

func TestDeleteUserSubject(t *testing.T) {
	Subject := domain.UserSubjects{"35", "BUSINESS", "THEORY", "FIRST"}
	result, err := DeleteUserSubject(Subject)
	assert.Nil(t, err)
	assert.True(t, result)

}
