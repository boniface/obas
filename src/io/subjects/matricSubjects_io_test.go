package io

import (
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/subjects"
	"testing"
)

func TestGetMatricSubjects(t *testing.T) {
	result, err := GetMatricSubjects()
	assert.Nil(t, err)

	assert.True(t, len(result) > 0)
}

func TestGetMatricSubject(t *testing.T) {
	expected := "THIDJ"
	result, err := GetMatricSubject("")
	assert.Nil(t, err)
	assert.Equal(t, expected, result.Description)

}

func TestCreateMatricSubject(t *testing.T) {
	sub := domain.MatricSubjects{"DS04", "MATH", "MATH", "FIRST"}
	result, err := CreateMatricSubject(sub)
	assert.Nil(t, err)
	assert.True(t, result)

}

func TestUpdateMatricSubject(t *testing.T) {
	sub := domain.MatricSubjects{"M023", "PHYSICS", "PHYSICS", "FIRST"}
	result, err := UpdateMatricSubject(sub)
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDeleteMatricSubject(t *testing.T) {
	sub := domain.MatricSubjects{"M023", "PHYSICS", "PHYSICS", "FIRST"}
	result, err := DeleteMatricSubject(sub)
	assert.Nil(t, err)
	assert.True(t, result)

}
