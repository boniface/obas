package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetApplicationResultes(t *testing.T) {
	valeu, _ := GetApplicationResultes()
	assert.NotNil(t, valeu)
	assert.Equal(t, valeu, "should be")
}
func TestGetApplicationResult(t *testing.T) {
	valeu, _ := GetApplicationResult("12153")
	assert.NotNil(t, valeu)
	assert.Equal(t, valeu, "should be")
}
func TestCreateApplicationResult(t *testing.T) {
	valeu, _ := CreateApplicationResult(0)
	assert.NotNil(t, valeu)
	assert.Equal(t, valeu, "should be")
}
func TestUpdateApplicationResult(t *testing.T) {
	valeu, _ := UpdateApplicationResult(0)
	assert.NotNil(t, valeu)
	assert.Equal(t, valeu, "should be")
}
func TestDeleteApplicationResult(t *testing.T) {
	valeu, _ := DeleteApplicationResult("121") // the id of the entity that we want to delete
	assert.NotNil(t, valeu)
	//assert.Equal(t,valeu,"should be")
}
