package application

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/application"
	"testing"
)

func TestGetApplicationResults(t *testing.T) {
	value, err := GetApplicationResults()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetApplicationResult(t *testing.T) {
	expected := "matric2019"
	value, err := GetApplicationResult("256")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.Date, expected)
}

func TestCreateApplicationResult(t *testing.T) {
	appType := domain.ApplicationResult{"256", "matric2019", "2019"}
	value, err := CreateApplicationResult(appType)
	assert.Nil(t, err)
	assert.True(t, value)
}
func TestUpdateApplicationResult(t *testing.T) {
	value, err := UpdateApplicationResult("")
	assert.Nil(t, err)
	assert.True(t, value)
}
func TestDeleteApplicationResult(t *testing.T) {
	appType := domain.ApplicationResult{"256", "matric2019", "2019"}
	result, err := DeleteApplicationResult(appType)
	assert.Nil(t, err)
	assert.True(t, result)
}
