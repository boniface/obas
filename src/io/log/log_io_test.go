package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogEvents(t *testing.T) {
	result, err := GetLogEvents()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, len(result) > 0)

}
