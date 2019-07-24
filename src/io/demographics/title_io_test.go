package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTitles(t *testing.T) {
	value, err := GetTitles()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetTitle(t *testing.T) {
	expected := ""
	value, err := GetTitle("")
	assert.Nil(t, err)
	assert.Equal(t, expected, value)
}

func TestCreateTitle(t *testing.T) {
	value, err := CreateTitle("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateTitle(t *testing.T) {
	value, err := UpdateTitle("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestDeleteTitle(t *testing.T) {
	value, err := DeleteTitle("")
	assert.Nil(t, err)
	assert.True(t, value)
}
