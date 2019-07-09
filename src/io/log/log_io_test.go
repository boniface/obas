package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogEvents(t *testing.T) {
	result, err := GetLogEvents()
	assert.Nil(t, err)

	assert.True(t, len(result) > 0)

}
