package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStudentDocuments(t *testing.T) {
	result, err := GetStudentDocuments()
	assert.Nil(t, err)

	assert.True(t, len(result) > 0)
}

func TestGetStudentDocument(t *testing.T) {
	expected := ""
	result, err := GetStudentDocument("")
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

}

func TestCreateStudentDocuments(t *testing.T) {
	result, err := CreateStudentDocuments("")
	assert.Nil(t, err)
	assert.True(t, result)

}

func TestUpdateStudentDocuments(t *testing.T) {
	result, err := UpdateStudentDocuments("")
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDeleteStudentDocuments(t *testing.T) {
	result, err := DeleteStudentDocuments("")
	assert.Nil(t, err)
	assert.True(t, result)

}
