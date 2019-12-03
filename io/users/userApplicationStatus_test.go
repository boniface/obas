package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
	"time"
)

var now = time.Now()

func TestCreateUserApplicationStatus(t *testing.T) {

	obj := domain.UserApplicationStatus{"000", "0002", now, "3939393"}
	resp, err := CreateUserApplicationStatus(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
func TestDeleteApplicationStatus(t *testing.T) {

	obj := domain.UserApplicationStatus{"000", "0002", now, "3939393"}
	resp, err := DeleteApplicationStatus(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
func TestGetUserApplicationStatues(t *testing.T) {
	resp, err := GetUserApplicationStatus("0003")
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
func TestGetUserApplicationStatus(t *testing.T) {
	resp, err := GetUserApplicationStatues()
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
func TestUpdateUserApplicationStatus(t *testing.T) {
	obj := domain.UserApplicationStatus{"000", "0002", now, "3939393"}
	resp, err := CreateUserApplicationStatus(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
