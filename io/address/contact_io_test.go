package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/address"
	"testing"
)

var entity = domain.ContactType{ContactTypeId: "1991", Name: "Christian M"}

func TestGetContacts(t *testing.T) {
	value, err := GetContactTypes()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)

}

func TestGetContactType(t *testing.T) {
	expected := entity
	value, err := GetContactType(entity.ContactTypeId)
	assert.Nil(t, err)
	assert.Equal(t, expected, value)
}

func TestCreateContactType(t *testing.T) {
	value, err := CreateContactType(entity)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateContactType(t *testing.T) {
	var expected = "Christian Muamba"
	var updated = domain.ContactType{ContactTypeId: "1991", Name: "Christian Muamba"}
	result, err := UpdateContactType(updated)
	assert.Nil(t, err)
	assert.True(t, result)
	value, err := GetContactType(entity.ContactTypeId)
	assert.Equal(t, expected, value.Name)

}
func TestDeleteContactType(t *testing.T) {
	value, err := DeleteContactType("")
	assert.Nil(t, err)
	assert.True(t, value)

}
