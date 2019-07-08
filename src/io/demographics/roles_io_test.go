package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRoles(t *testing.T) {
	value, _ := GetRoles()
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestGetRole(t *testing.T) {
	value, _ := GetRole("121")
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestCreateRole(t *testing.T) {
	value, _ := CreateRole(0)
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestUpdateRole(t *testing.T) {
	value, _ := UpdateRole(0)
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}

func TestDeleteRole(t *testing.T) {
	value, _ := DeleteRole("121")
	assert.Equal(t, value, "expected here")
	assert.NotNil(t, value)
}
