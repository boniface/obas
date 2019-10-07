package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

func TestGetUserPasswords(t *testing.T) {
	value, err := GetUserPasswords()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetUserPassword(t *testing.T) {
	expected := "H7I@#ioP"
	value, err := GetUserPassword("57")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.Password, expected)
}

func TestUpdateUserPassword(t *testing.T) {
	uPassword := domain.UserPassword{"51", "U4B6ER@D#"}
	value, err := UpdateUserPassword(uPassword)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestCreateUserPassword(t *testing.T) {
	uPassword := domain.UserPassword{"35", "123_@Pass"}
	value, err := CreateUserPassword(uPassword)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestDeleteUserContact(t *testing.T) {
	uPassword := domain.UserPassword{"35", "123_@Pass"}
	value, err := DeleteUserPassword(uPassword)
	assert.Nil(t, err)
	assert.True(t, value)
}
