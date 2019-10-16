package security

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestgetApiKey(t *testing.T) {
	expected := "WC"
	result, err := getApiKey("")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, result.Value, expected)
}
