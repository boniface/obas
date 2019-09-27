package institutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/institutions"
	"testing"
)

var sch = domain.School{"98", "CAPETOWN HIGH", "CBD", "WC"}

func TestGetSchools(t *testing.T) {
	value, err := GetSchools()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetSchool(t *testing.T) {
	expected := sch
	value, err := GetSchool(sch.SchoolId)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value, expected)
}

func TestCreateSchool(t *testing.T) {
	value, err := CreateSchool(sch)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateSchool(t *testing.T) {
	var expected = "NGP HIGH"
	var updated = domain.School{"87", "NGP HIGH", "DFG", "WC"}
	result, err := UpdateSchool(updated)
	assert.Nil(t, err)
	assert.True(t, result)
	value, err := GetSchool(sch.SchoolId)
	assert.Equal(t, expected, value.SchoolName)
}

func TestDeleteSchool(t *testing.T) {
	value, err := DeleteSchool(sch)
	assert.Nil(t, err)
	assert.True(t, value)
}
