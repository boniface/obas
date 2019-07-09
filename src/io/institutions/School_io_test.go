package institutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSchools(t *testing.T) {
	value, err := GetSchools()
	assert.Nil(t, err)
	assert.Equal(t, value, "entity", "Return entity")
}

func TestGetSchool(t *testing.T) {
	value, err := GetSchool("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestCreateSchool(t *testing.T) {
	value, err := CreateSchool("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestUpdateDocument(t *testing.T) {
	value, err := UpdateSchool("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}

func TestDeleteDocument(t *testing.T) {
	value, err := DeleteSchool("")
	assert.NotNil(t, err)
	assert.Equal(t, value, "Return entity")
}
