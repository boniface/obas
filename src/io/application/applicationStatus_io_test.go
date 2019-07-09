package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetApplicationStatuses(t *testing.T) {
	value, _ := GetApplicationResultes()
	assert.NotNil(t, value)
	assert.Equal(t, value, "expected value")

}
func TestGetApplicationStatus(t *testing.T) {
	value, _ := GetApplicationResult("1222")
	assert.NotNil(t, value)
	assert.Equal(t, value, "expected value")
}
func TestCreateApplicationStatus(t *testing.T) {
	value, _ := CreateApplicationStatus(0)
	assert.NotNil(t, value)
	assert.Equal(t, value, "expected value")
}
func TestUpdateApplicationStatus(t *testing.T) {
	value, _ := UpdateApplicationStatus(0)
	assert.NotNil(t, value)
	assert.Equal(t, value, "expected value")
}
func TestDeleteApplicationStatus(t *testing.T) {
	value, _ := DeleteApplicationStatus("1231")
	assert.NotNil(t, value)
	assert.Equal(t, value, "expected value")
}
