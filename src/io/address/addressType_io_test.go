package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAddress(t *testing.T) {
	value, err := CreateAddress("")
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
	value, err := GetAddresses()
	assert.Nil(t, err)
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
