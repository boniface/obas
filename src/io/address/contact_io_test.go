package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/address"
	"testing"
)

func TestGetContacts(t *testing.T) {

	value, err := GetContactTypes()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)

}

func TestGetContactType(t *testing.T) {
	expected := ""
	value, err := GetContactType("")
	assert.Nil(t, err)
	assert.Equal(t, expected, value)
}

func TestCreateContactType(t *testing.T) {
	contType := domain.ContactType{"CONTACT", "CONTACT ADDRESS"}
	value, err := CreateContactType(contType)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateContactType(t *testing.T) {
	value, err := UpdateContactType("")
	assert.Nil(t, err)
	assert.True(t, value)
}
func TestDeleteContactType(t *testing.T) {
	value, err := DeleteContactType("")
	assert.Nil(t, err)
	assert.True(t, value)

}
