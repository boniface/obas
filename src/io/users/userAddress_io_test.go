package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/users"
	"testing"
)

func TestGetUserAddresses(t *testing.T) {
	value, err := GetUserAddresses()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetUserAddress(t *testing.T) {
	expected := "136 BREE ST"
	value, err := GetUserAddress("15")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.PhysicalAddress, expected)
}

func TestCreateUserAddress(t *testing.T) {
	uAddr := domain.UserAddress{"15", "136 BREE ST", "7894"}
	value, err := CreateUserAddress(uAddr)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateUserAddress(t *testing.T) {
	uAddr := domain.UserAddress{"16", "136 BREE ST", "7894"}
	value, err := UpdateUserAddress(uAddr)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteUserAddress(t *testing.T) {
	uAddr := domain.UserAddress{"15", "136 BREE ST", "7894"}
	value, err := DeleteUserAddress(uAddr)
	assert.Nil(t, err)
	assert.True(t, value)
}
