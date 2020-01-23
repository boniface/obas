package demographics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/demographics"
	"testing"
)

func TestGetTitles(t *testing.T) {
	value, err := GetTitles()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetTitle(t *testing.T) {
	expected := "DR"
	value, err := GetTitle("")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, expected, value.TitleName)
}

func TestCreateTitle(t *testing.T) {
	title := domain.Title{"", "Mrs"}
	value, err := CreateTitle(title)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestUpdateTitle(t *testing.T) {
	title := domain.Title{"47", "SIR"}
	value, err := UpdateTitle(title)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteTitle(t *testing.T) {
	title := domain.Title{"47", "SIR"}
	value, err := DeleteTitle(title)
	assert.Nil(t, err)
	assert.True(t, value)
}
