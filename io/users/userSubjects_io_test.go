package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

var Subject = domain.UserSubjects{"56", "INFORMATION", "THEORY", "FIRST"}

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
	assert.Equal(t, expected, result.UserId)

}

func TestCreateUserSubject(t *testing.T) {
	result, err := CreateUserSubject(Subject)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)

}

func TestUpdateUserSubject(t *testing.T) {
	var expected = "SOFTWARE"
	var updated = domain.UserSubjects{"40", "SOFTWARE", "UML DIAGRAM", "SECOND"}
	result, err := UpdateUserSubject(updated)
	assert.Nil(t, err)
	assert.True(t, result)
	value, err := GetUserSubject(Subject.UserId)
	assert.Equal(t, expected, value.UserId)
}

func TestDeleteUserSubject(t *testing.T) {
	Subject := domain.UserSubjects{"35", "BUSINESS", "THEORY", "FIRST"}
	result, err := DeleteUserSubject(Subject)
	assert.Nil(t, err)
	assert.True(t, result)

}
