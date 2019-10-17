package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

func TestGetUserContacts(t *testing.T) {
	result, err := GetUserContacts()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, len(result) > 0)
}

func TestGetUserContact(t *testing.T) {
	expected := "0838956987"
	result, err := GetUserContact("52")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, expected, result.Contact)

}

func TestCreateUserContact(t *testing.T) {
	usrContact := domain.UserContacts{"895675624", "78965412", "m@m.com"}
	result, err := CreateUserContact(usrContact)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)

}

func TestUpdateUserContact(t *testing.T) {
	usrContact := domain.UserContacts{"78965412", "8596932", "m@g.com"}
	result, err := UpdateUserContact(usrContact)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)
}

func TestDeleteUserContacts(t *testing.T) {
	usrContact := domain.UserContacts{"78965412", "8596932", "m@g.com"}
	result, err := DeleteUserContact(usrContact)
	assert.Nil(t, err)
	assert.True(t, result)

}
