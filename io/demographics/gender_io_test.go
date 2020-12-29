package demographics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/demographics"
	"testing"
)

func TestGetGenders(t *testing.T) {
	value, err := GetGenders()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, len(value) > 0)
}

func TestGetGender(t *testing.T) {
	expected := "MALE"
	value, err := GetGender("5")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.GenderName, expected)
}

func TestCreateGender(t *testing.T) {
	gend := domain.Gender{"", "Male"}
	value, err := CreateGender(gend)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateGender(t *testing.T) {
	gend := domain.Gender{"4", "FEMALE"}
	value, err := UpdateGender(gend)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestDeleteGender(t *testing.T) {
	gend := domain.Gender{"4", "FEMALE"}
	value, err := DeleteGender(gend)
	assert.Nil(t, err)
	assert.True(t, value)
}
