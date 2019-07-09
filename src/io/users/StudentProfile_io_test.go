package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStudentProfiles(t *testing.T) {
	result, err := GetStudentProfiles()
	assert.Nil(t, err)

	assert.True(t, len(result) > 0)
}

func TestGetStudentProfile(t *testing.T) {
	expected := ""
	result, err := GetStudentProfile("")
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

}

func TestCreateStudentProfiles(t *testing.T) {
	result, err := CreateStudentProfiles("")
	assert.Nil(t, err)
	assert.True(t, result)

}

func TestUpdateStudentProfiles(t *testing.T) {
	result, err := UpdateStudentProfiles("")
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDeleteStudentProfiles(t *testing.T) {
	result, err := DeleteStudentProfiles("")
	assert.Nil(t, err)
	assert.True(t, result)

}
