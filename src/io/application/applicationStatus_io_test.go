package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetApplicationStatuses(t *testing.T) {
	value, err := GetApplicationResults()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)

}

func TestGetApplicationStatus(t *testing.T) {
	expected := ""
	value, err := GetApplicationResult("")
	assert.Nil(t, err)
	assert.Equal(t, value, expected)
}

func TestCreateApplicationStatus(t *testing.T) {
	value, err := CreateApplicationStatus("")
	assert.Nil(t, err)
	assert.True(t, value)
}
func TestUpdateApplicationStatus(t *testing.T) {
	value, err := UpdateApplicationStatus("")
	assert.Nil(t, err)
	assert.True(t, value)
}
func TestDeleteApplicationStatus(t *testing.T) {
	value, err := DeleteApplicationStatus("")
	assert.Nil(t, err)
	assert.True(t, value)
}
