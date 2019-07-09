package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStudentResults(t *testing.T) {
	result, err := GetStudentResults()
	assert.Nil(t, err)

	assert.True(t, len(result) > 0)
}

func TestGetStudentResult(t *testing.T) {
	expected := ""
	result, err := GetStudentResult("")
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

}

func TestCreateStudentResults(t *testing.T) {
	result, err := CreateStudentResults("")
	assert.Nil(t, err)
	assert.True(t, result)

}

func TestUpdateStudentResults(t *testing.T) {
	result, err := UpdateStudentResults("")
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDeleteStudentResults(t *testing.T) {
	result, err := DeleteStudentResults("")
	assert.Nil(t, err)
	assert.True(t, result)

}
