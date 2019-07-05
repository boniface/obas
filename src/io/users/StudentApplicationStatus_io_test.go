package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStudentApplicationStatuses2(t *testing.T) {
	result, err := GetProcessingStatusTypes()
	assert.Nil(t, err)

	assert.True(t, len(result) > 0)
}

func TestGetStudentApplicationStatuse(t *testing.T) {
	expected := ""
	result, err := GetStudentApplicationStatus("")
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

}

func TestCreateStudentApplicationStatus(t *testing.T) {
	result, err := CreateStudentApplicationStatus("")
	assert.Nil(t, err)
	assert.True(t, result)

}

func TestUpdateStudentApplicationStatus(t *testing.T) {
	result, err := UpdateStudentApplicationStatus("")
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDeleteStudentApplicationStatus(t *testing.T) {
	result, err := DeleteStudentApplicationStatus("")
	assert.Nil(t, err)
	assert.True(t, result)

}
