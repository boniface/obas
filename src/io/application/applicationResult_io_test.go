package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetApplicationResults(t *testing.T) {
	value, err := GetApplicationResults()
	assert.Nil(t, err)
	assert.True(t, len(value) > 0)
}

func TestGetApplicationResult(t *testing.T) {
	expected := ""
	value, err := GetApplicationResult("")
	assert.Nil(t, err)
	assert.Equal(t, value, expected)
}

func TestCreateApplicationResult(t *testing.T) {

	value, err := CreateApplicationResult("")
	assert.Nil(t, err)
	assert.True(t, value)
}
func TestUpdateApplicationResult(t *testing.T) {
	value, err := UpdateApplicationResult("")
	assert.Nil(t, err)
	assert.True(t, value)
}
func TestDeleteApplicationResult(t *testing.T) {
	result, err := DeleteApplicationResult("")
	assert.Nil(t, err)
	assert.True(t, result)
}
