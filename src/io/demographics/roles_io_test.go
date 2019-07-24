package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRoles(t *testing.T) {
	value, err := GetRoles()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetRole(t *testing.T) {
	expected := ""
	value, err := GetRole("")
	assert.Nil(t, err)
	assert.Equal(t, expected, value)
}

func TestCreateRole(t *testing.T) {
	value, err := CreateRole("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateRole(t *testing.T) {
	value, err := UpdateRole("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestDeleteRole(t *testing.T) {
	value, err := DeleteRole("")
	assert.Nil(t, err)
	assert.True(t, value)
}
