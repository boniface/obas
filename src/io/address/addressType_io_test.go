package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAddressType(t *testing.T) {
	addrType := Address{"SNAPCHAT", "SNAPCHAT ADDRESS"}
	value, err := CreateAddressType(addrType)
	assert.Nil(t, err)
	assert.True(t, value)
}
func TestGetAddressType(t *testing.T) {
	expected := ""
	value, err := GetAddressType("")
	assert.Nil(t, err)
	assert.Equal(t, value, expected)

}
func TestGetAddressTypes(t *testing.T) {
	value, err := GetAddresses()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestUpdateAddressType(t *testing.T) {
	value, err := UpdateAddressType("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestDeleteAddressType(t *testing.T) {
	value, err := DeleteAddressType("")
	assert.Nil(t, err)
	assert.True(t, value)
}