package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/application"
	"testing"
)

func TestGetApplicationStatuses(t *testing.T) {
	value, err := GetApplicationResults()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)

}

func TestGetApplicationStatus(t *testing.T) {
	expected := "Pending"
	value, err := GetApplicationResult("212")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.Date, expected)
}

func TestCreateApplicationStatus(t *testing.T) {
	appType := domain.ApplicationStatus{"212", "Pending", "2020"}
	value, err := CreateApplicationStatus(appType)
	assert.Nil(t, err)
	assert.True(t, value)
}
func TestUpdateApplicationStatus(t *testing.T) {
	appType := domain.ApplicationStatus{"212", "Pending", "2020"}
	value, err := UpdateApplicationStatus(appType)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}
func TestDeleteApplicationStatus(t *testing.T) {
	appType := domain.ApplicationStatus{"212", "Pending", "2020"}
	value, err := DeleteApplicationStatus(appType)
	assert.Nil(t, err)
	assert.True(t, value)
}
