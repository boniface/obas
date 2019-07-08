package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMatricSubjects(t *testing.T) {
	result, err := GetMatricSubjects()
	assert.Nil(t, err)

	assert.True(t, len(result) > 0)
}

func TestGetMatricSubject(t *testing.T) {
	expected := ""
	result, err := GetMatricSubject("")
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

}

func TestCreateMatricSubject(t *testing.T) {
	result, err := CreateMatricSubject("")
	assert.Nil(t, err)
	assert.True(t, result)

}

func TestUpdateMatricSubject(t *testing.T) {
	result, err := UpdateMatricSubject("")
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDeleteMatricSubject(t *testing.T) {
	result, err := DeleteMatricSubject("")
	assert.Nil(t, err)
	assert.True(t, result)

}
