package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProcessingStatusTypes(t *testing.T) {
	result, err := GetProcessingStatusTypes()
	assert.Nil(t, err)

	assert.True(t, len(result) > 0)
}

func TestGetProcessingStatusType(t *testing.T) {
	expected := ""
	result, err := GetProcessingStatusType("")
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

}

func TestCreateProcessingStatusType(t *testing.T) {
	result, err := CreateAdmin("")
	assert.Nil(t, err)
	assert.True(t, result)

}

func TestUpdateProcessingStatusType(t *testing.T) {
	result, err := UpdateProcessingStatusType("")
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDeleteProcessingStatustype(t *testing.T) {
	result, err := DeleteProcessingStatustype("")
	assert.Nil(t, err)
	assert.True(t, result)

}
