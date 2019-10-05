package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

func TestGetUserIntstitutions(t *testing.T) {
	value, err := GetUserIntstitutions()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetUserIntstitution(t *testing.T) {
	expected := "CUT"
	value, err := GetUserIntstitution("87")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.Name, expected)
}

func TestCreateUserIntstitution(t *testing.T) {
	uInt := domain.UserInstitution{"86", "UNISA"}
	value, err := CreateUserIntstitution(uInt)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestUpdateUserIntstitution(t *testing.T) {
	uInt := domain.UserInstitution{"86", "UNISA"}
	value, err := DeleteUserIntstitution(uInt)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteUserIntstitution(t *testing.T) {
	uInt := domain.UserInstitution{"86", "UNISA"}
	value, err := DeleteUserIntstitution(uInt)
	assert.Nil(t, err)
	assert.True(t, value)
}
