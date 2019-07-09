package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAddress(t *testing.T) {
	value, nulvalue := CreateAddress(0) //we should specify the entity
	assert.Equal(t, value, "address destails here", "should be equal")
	assert.Nil(t, nulvalue, nil, "this should be null")
}
func TestGetAddress(t *testing.T) {
	value, nulvalue := GetAddress("1212") //we should specify the id here
	assert.Equal(t, "should return the entity", "should return the entity")
	assert.NotNil(t, value)   // testing the null return
	assert.Error(t, nulvalue) // if we want to test in case of an error
}
func TestGetAddresses(t *testing.T) {
	value, nullvalue := GetAddresses()
	assert.Equal(t, value, "should contain the expected values here")
	assert.NotNil(t, value)  // if we want to make sure that this value doesnt return null when is shouldnt
	assert.Nil(t, nullvalue) // if we want to make sure that this is returning null when we putting wrong valeus
}
func TestUpdateAddress(t *testing.T) {
	value, nullvalue := UpdateAddress(0) //we should specify the entity
	assert.Equal(t, value, "should contain the expected values here")
	assert.NotNil(t, value)  // if we want to make sure that this value doesnt return null when is shouldnt
	assert.Nil(t, nullvalue) // if we want to make sure that this is returning null when we putting wrong valeus
}
func TestDeleteAddress(t *testing.T) {
	value, nullvalue := DeleteAddress(0) //we should specify the entity to delete
	//assert.Equal(t,value,"should contain the expected values here")
	assert.NotNil(t, value)  // if we want to make sure that this value doesnt return null when is shouldnt
	assert.Nil(t, nullvalue) // if we want to make sure that this is returning null when we putting wrong valeus
}
