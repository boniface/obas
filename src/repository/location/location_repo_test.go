package location

import (
	"fmt"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	CreateAccount()
}

func TestReadAccount(t *testing.T) {
	account := ReadAccount()
	fmt.Println(" the Account is ", account)
}
