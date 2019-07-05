package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStudentContacts(t *testing.T) {
	result, err := GetStudentContacts()
	assert.Nil(t, err)

	assert.True(t, len(result) > 0)
}

func TestGetStudentContact(t *testing.T) {
	expected := ""
	result, err := GetStudentContact("")
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

}

func TestCreateStudentContact(t *testing.T) {
	result, err := CreateStudentContact("")
	assert.Nil(t, err)
	assert.True(t, result)

}

func TestUpdateStudentContact(t *testing.T) {
	result, err := UpdateStudentContact("")
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDeleteStudentContact(t *testing.T) {
	result, err := DeleteStudentContact("")
	assert.Nil(t, err)
	assert.True(t, result)

}
