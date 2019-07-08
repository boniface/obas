package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetApplicationtype(t *testing.T) {
	value, _ := GetApplicationtype("121")
	assert.NotNil(t, value)
	assert.Equal(t, value, "expected entity value")
}
func TestGetApplicationTypes(t *testing.T) {
	value, _ := GetApplicationTypes()
	assert.NotNil(t, value)
	assert.Equal(t, value, "expected entity value")
}
func TestCreateApplicationtype(t *testing.T) {
	value, _ := GetApplicationtype("121")
	assert.NotNil(t, value)
	assert.Equal(t, value, "expected entity value")
}
func TestUpdateApplication(t *testing.T) {
	value, _ := UpdateApplicationtype(0)
	assert.NotNil(t, value)
	assert.Equal(t, value, "expected entity value")
}
func TestDeleteApplicationtype(t *testing.T) {
	value, _ := DeleteApplicationtype("121")
	assert.NotNil(t, value)
	assert.Equal(t, value, "expected entity value")
}
