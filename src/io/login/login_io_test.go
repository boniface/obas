package login

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoRegister(t *testing.T) {
	result, err := DoRegister("test@gmail.com")
	if err != nil {
		fmt.Println("This has errors: ", err.Error())
	}

	assert.True(t, result)
}