package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUniversityCourses(t *testing.T) {
	result, err := GetUniversityCourses()
	assert.Nil(t, err)

	assert.True(t, len(result) > 0)
}

func TestGetUniversityCourse(t *testing.T) {
	expected := ""
	result, err := GetUniversityCourse("")
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

}

func TestCreateUniversityCourses(t *testing.T) {
	result, err := CreateUniversityCourses("")
	assert.Nil(t, err)
	assert.True(t, result)

}

func TestUpdateUniversityCourses(t *testing.T) {
	result, err := UpdateUniversityCourses("")
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDeleteUniversityCourses(t *testing.T) {
	result, err := DeleteUniversityCourses("")
	assert.Nil(t, err)
	assert.True(t, result)

}
