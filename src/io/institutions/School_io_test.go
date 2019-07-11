package institutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSchools(t *testing.T) {
	value, err := GetSchools()
	assert.Nil(t, err)
	assert.True(t, len(value) > 0)
}

func TestGetSchool(t *testing.T) {
	expected := ""
	value, err := GetSchool("")
	assert.Nil(t, err)
	assert.Equal(t, value, expected)
}

func TestCreateSchool(t *testing.T) {
	value, err := CreateSchool("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateSchool(t *testing.T) {
	value, err := UpdateSchool("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestDeleteSchool(t *testing.T) {
	value, err := DeleteSchool("")
	assert.Nil(t, err)
	assert.True(t, value)
}
