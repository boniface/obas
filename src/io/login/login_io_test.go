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
func TestLogin_io(t *testing.T) {
	result, err := Login_io("espoirditekemena@gmail.com", "qNgXgA9I")
	if err != nil {
		fmt.Println("This has errors: ", err.Error())
	}
	fmt.Println(" thre Result is ", err)
	assert.NotNil(t, result)
}
