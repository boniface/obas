package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/subjects"
	"testing"
)

func TestGetMatricSubjects(t *testing.T) {
	result, err := GetMatricSubjects()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, len(result) > 0)
}

func TestGetMatricSubject(t *testing.T) {
	expected := "Bachelor"
	result, err := GetMatricSubject("123")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, expected, result.Name)

}

func TestCreateMatricSubject(t *testing.T) {
	sub := domain.MatricSubjects{"DS04", "MATH", "MATH", "FIRST"}
	result, err := CreateMatricSubject(sub)
	assert.Nil(t, err)
	assert.True(t, result)

}

func TestUpdateMatricSubject(t *testing.T) {
	sub := domain.MatricSubjects{"DS04", "CALCULUS", "MATH", "FIRST"}
	result, err := UpdateMatricSubject(sub)
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDeleteMatricSubject(t *testing.T) {
	sub := domain.MatricSubjects{"DS04", "MATH", "MATH", "FIRST"}
	result, err := DeleteMatricSubject(sub)
	assert.Nil(t, err)
	assert.True(t, result)

}
