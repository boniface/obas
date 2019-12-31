package applications

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetApplicationTypes(t *testing.T) {
	value, err := GetApplicationTypes()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}
func TestGetApplicationtype(t *testing.T) {
	expected := ""
	value, err := GetApplicationType("")
	assert.Nil(t, err)
	assert.Equal(t, expected, value)
}

func TestCreateApplicationtype(t *testing.T) {
	value, err := CreateApplicationType("")
	assert.Nil(t, err)
	assert.True(t, value)
}
func TestUpdateApplication(t *testing.T) {
	value, err := UpdateApplicationType("")
	assert.Nil(t, err)
	assert.True(t, value)
}
func TestDeleteApplicationtype(t *testing.T) {
	value, err := DeleteApplicationType("")
	assert.Nil(t, err)
	assert.True(t, value)
}
