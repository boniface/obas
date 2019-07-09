package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStudentDemographics(t *testing.T) {
	result, err := GetStudentDemographics()
	assert.Nil(t, err)

	assert.True(t, len(result) > 0)
}

func TestGetStudentDemographic(t *testing.T) {
	expected := ""
	result, err := GetStudentDemographic("")
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

}

func TestCreateStudentDemographics(t *testing.T) {
	result, err := CreateStudentDemographics("")
	assert.Nil(t, err)
	assert.True(t, result)

}

func TestUpdateStudentDemographics(t *testing.T) {
	result, err := UpdateStudentDemographics("")
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDeleteStudentDemographics(t *testing.T) {
	result, err := DeleteStudentDemographics("")
	assert.Nil(t, err)
	assert.True(t, result)

}
