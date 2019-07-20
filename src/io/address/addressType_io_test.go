package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/address"
	"testing"
)

func TestCreateAddress(t *testing.T) {
	addrType := domain.AddressType{"EMAIL", "EMAIL ADDRESS"}
	value, err := CreateAddressType(addrType)
	assert.Nil(t, err)
	assert.True(t, value)
}
func TestGetAddress(t *testing.T) {
	expected := ""
	value, err := GetAddress("") //we should specify the id here
	assert.Nil(t, err)
	assert.Equal(t, value, expected)

}
func TestGetAddresses(t *testing.T) {
	value, err := GetAddressTypes()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestUpdateAddress(t *testing.T) {
	value, err := UpdateAddress("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestDeleteAddress(t *testing.T) {
	value, err := DeleteAddress("")
	assert.Nil(t, err)
	assert.True(t, value)
}
