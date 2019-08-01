package institutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/institutions"
	"testing"
)

var univ = domain.University{"14", "DUT", "TECHNO", "EC"}

func TestGetUniversitys(t *testing.T) {
	value, err := GetUniversitys()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetUniversity(t *testing.T) {
	expected := univ
	value, err := GetUniversity(univ.UniversityId)
	assert.Nil(t, err)
	assert.Equal(t, value, expected)
}

func TestCreateUniversity(t *testing.T) {
	value, err := CreateUniversity(univ)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateUniversity(t *testing.T) {
	var expected = "DUT"
	var univ = domain.University{"14", "DUT", "ALL", "KZN"}
	result, err := UpdateUniversity(univ)
	assert.Nil(t, err)
	assert.True(t, result)
	value, err := GetUniversity(univ.UniversityId)
	assert.Equal(t, expected, value.UniversityName)
}

func TestDeleteUniversity(t *testing.T) {
	value, err := DeleteUniversity(univ)
	assert.Nil(t, err)
	assert.True(t, value)
}
