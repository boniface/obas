package institutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUniversitys(t *testing.T) {
	value, err := GetUniversitys()
	assert.Nil(t, err)
	assert.Equal(t, value, "entity", "Return entity")
}

func TestGetUniversity(t *testing.T) {
	value, err := GetUniversity("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestCreateUniversity(t *testing.T) {
	value, err := CreateUniversity("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestUpdateUniversity(t *testing.T) {
	value, err := UpdateUniversity("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestDeleteUniversity(t *testing.T) {
	value, err := DeleteUniversity("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}
