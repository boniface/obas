package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/subjects"
	"testing"
)

func TestGetUniversityCourses(t *testing.T) {
	result, err := GetUniversityCourses()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, len(result) > 0)
}

func TestGetUniversityCourse(t *testing.T) {
	expected := "UCT"
	result, err := GetUniversityCourse("ADS04")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, expected, result.Name)

}

func TestCreateUniversityCourses(t *testing.T) {
	result := domain.UniversityCourses{"ADS04", "3YEAR", "UCT", "FULL", "THIRD"}
	value, err := CreateUniversityCourses(result)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)

}

func TestUpdateUniversityCourses(t *testing.T) {
	result := domain.UniversityCourses{"ADS04", "3YEAR", "UCT", "FULL", "THIRD"}
	value, err := UpdateUniversityCourses(result)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteUniversityCourses(t *testing.T) {
	result := domain.UniversityCourses{"ADS04", "3YEAR", "UCT", "FULL", "THIRD"}
	value, err := DeleteUniversityCourses(result)
	assert.Nil(t, err)
	assert.True(t, value)

}
