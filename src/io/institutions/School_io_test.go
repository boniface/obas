package institutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/institutions"
	"testing"
)

func TestGetSchools(t *testing.T) {
	value, err := GetSchools()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetSchool(t *testing.T) {
	expected := "NGP HIGH"
	value, err := GetSchool("89")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value, expected)
}

func TestCreateSchool(t *testing.T) {
	sch := domain.School{"89", "NGP HIGH", "DFG", "WC"}
	value, err := CreateSchool(sch)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateSchool(t *testing.T) {
	sch := domain.School{"87", "NGP HIGH", "DFG", "WC"}
	value, err := UpdateSchool(sch)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteSchool(t *testing.T) {
	sch := domain.School{"87", "NGP HIGH", "DFG", "WC"}
	value, err := DeleteSchool(sch)
	assert.Nil(t, err)
	assert.True(t, value)
}
