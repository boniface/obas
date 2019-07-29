package institutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/institutions"
	"testing"
)

func TestGetUniversitys(t *testing.T) {
	value, err := GetUniversitys()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetUniversity(t *testing.T) {
	expected := "DUT"
	value, err := GetUniversity("14")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value, expected)
}

func TestCreateUniversity(t *testing.T) {
	univ := domain.University{"14", "DUT", "TECHNO", "EC"}
	value, err := CreateUniversity(univ)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateUniversity(t *testing.T) {
	univ := domain.University{"14", "DUT", "TECHNO", "KZN"}
	value, err := UpdateUniversity(univ)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestDeleteUniversity(t *testing.T) {
	univ := domain.University{"14", "DUT", "TECHNO", "KZN"}
	value, err := DeleteUniversity(univ)
	assert.Nil(t, err)
	assert.True(t, value)
}
