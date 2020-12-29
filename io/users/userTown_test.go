package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

type UserTown2 domain.UserTown

var token = "eyJraWQiOiJvYmFzYXBpX29uXzE1NS4yMzguMzIuMjE5IiwiYWxnIjoiRVMyNTYifQ.eyJpc3MiOiJIQVNIQ09ERS5aTSIsImF1ZCI6IlNJVEVVU0VSUyIsImV4cCI6MTU3OTc2MzcyNiwianRpIjoiOXVHcnVSWF9yUy1QZ09URWJIM0psUSIsImlhdCI6MTU3OTY3NzMyNiwibmJmIjoxNTc5Njc3MjA2LCJzdWIiOiJTaXRlIEFjY2VzcyIsImVtYWlsIjoiZXNwb2lyZGl0ZWtlbWVuYUBnbWFpbC5jb20iLCJyb2xlIjoiU1RSMDAxIn0.Wt9HH1YwbhHXH9EmJTfEFeMCCyYIx5i4Slh3Mb6cHdHgu3LyJ99bgKiIzfuzEbtden4DNJ7K1PDWEF0b7325Iw"

func TestCreateUserTown(t *testing.T) {
	//obj :=domain.UserTown{"0908898","980989"}
	//result,err:=CreateUserTown(obj)

}
func TestUpdateUserTown(t *testing.T) {
	obj := domain.UserTown{"espoirditekemena@gmail.com", "BOTU-8UWZZ"}
	value, err := UpdateUserTown(obj, token)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value)
}
func TestGetUserTown(t *testing.T) {
	value, err := GetUserTown("espoirditekemena@gmail.com")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value)
}
