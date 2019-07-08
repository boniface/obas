package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTitles(t *testing.T) {
	value, _ := GetTitles()
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestGetTitle(t *testing.T) {
	value, _ := GetTitle("121")
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestCreateTitle(t *testing.T) {
	value, _ := CreateTitle(0)
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestUpdateTitle(t *testing.T) {
	value, _ := UpdateTitle(0)
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestDeleteTitle(t *testing.T) {
	value, _ := DeleteTitle(0)
	//assert.Equal(t,value,"expected here")
	//assert.NotNil(t,value)
	assert.True(t, value)
}
