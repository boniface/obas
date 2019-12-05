package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/util"
	"testing"
)

func TestGetStatuses(t *testing.T) {
	result, err := GetStatuses()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, len(result) > 0)
}

func TestCreateStatus(t *testing.T) {
	entity := domain.GenericStatus{"", "Rejected", ""}
	result, err := CreateStatus(entity)
	assert.Nil(t, err)
	fmt.Println(" The Result", result)
}