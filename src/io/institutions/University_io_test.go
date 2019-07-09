package institutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUniversitys(t *testing.T) {
	value, err := GetUniversitys()
	assert.Nil(t, err)
	assert.True(t, len(value) > 0)
}

func TestGetUniversity(t *testing.T) {
	expected := ""
	value, err := GetUniversity("")
	assert.Nil(t, err)
	assert.Equal(t, value, expected)
}

func TestCreateUniversity(t *testing.T) {
	value, err := CreateUniversity("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateUniversity(t *testing.T) {
	value, err := UpdateUniversity("")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestDeleteUniversity(t *testing.T) {
	value, err := DeleteUniversity("")
	assert.Nil(t, err)
	assert.True(t, value)
}
