package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Contact struct {
	phone  int
	street string
	suburb string
}

func TestGetContacts(t *testing.T) {
	value, mynull := GetContactTypes()
	assert.Equal(t, value, "all the address should be returned", "they should be true")
	assert.Nil(t, mynull, nil)
	//assert := assert.New(t)

	// assert equality
	//assert.Equal(122, 123, "they should be equal")

}
func TestCreateContactType(t *testing.T) {
	//p:=Contact{2344,"rebeick","goodwood"}
	assert.True(t, false, "they should be true")

	//contacttype:= CreateContactType(p.phone,p.street,p.suburb)
}
func TestGetContactType(t *testing.T) {
	mycontact, _ := GetContactType("0011")
	assert.Equal(t, mycontact, "ewewe", "they should be true")
}
func TestUpdateContactType(t *testing.T) {
	value, _ := UpdateContactType(0) // we should put the entity valeus
	assert.Equal(t, value, 0, "this should be equal")

}
func TestDeleteContactType(t *testing.T) {

	valeu, nullvalue := DeleteContactType(0)
	assert.True(t, valeu, true, "this should be true")
	assert.Nil(t, nullvalue, nil, "this should be null")

}
