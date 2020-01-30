package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

//var token ="eyJraWQiOiJURVNUX1BIUkFTRSIsImFsZyI6IkVTMjU2In0.eyJpc3MiOiJIQVNIQ09ERS5aTSIsImF1ZCI6IlNJVEVVU0VSUyIsImV4cCI6MTU4MDIxNTU2OCwianRpIjoiX284WGVZX1h3WGNVU2pTWjZMTGZkdyIsImlhdCI6MTU4MDEyOTE2OCwibmJmIjoxNTgwMTI5MDQ4LCJzdWIiOiJTaXRlIEFjY2VzcyIsImVtYWlsIjoiZXNwb2lyZGl0ZWtlbWVuYUBnbWFpbC5jb20iLCJyb2xlIjoiQUFJSS05Q1pEViJ9.2FCbuRUZbFygGDD7KoGiEpYlIWhgz6b2IZ8_n1x3m3NObL47eLn6uFbkCxy26UPkA-RH3ylcDJeHPBltd3w8MA"

func TestGetUserRoles(t *testing.T) {
	result, err := GetUserRoles()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}

func TestGetUserRole(t *testing.T) {
	///expected := "25"
	result, err := GetUserRole("STR001")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	//assert.Equal(t, expected, result.RoleId)

}

func TestCreateUserRole(t *testing.T) {
	role := domain.UserRole{"35", "ADMIN"}
	result, err := CreateUserRole(role)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)

}

func TestUpdateUserRole(t *testing.T) {
	role := domain.UserRole{"216093805@mycput.ac.za", "MMMMM"}
	result, err := UpdateUserRole(role, token)
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
