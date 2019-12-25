package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

func TestCreateUserTertiaryCourse(t *testing.T) {
	obj := domain.UserTertiaryCourse{}
	result, err := CreateUserTertiaryCourse(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetUserTertiaryCourse(t *testing.T) {
	result, err := GetUserTertiaryCourse("")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetUserApplicationWithAppId(t *testing.T) {
	result, err := GetUserApplicationWithAppId("")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetUserTertiaryCourseForApp(t *testing.T) {
	result, err := GetUserTertiaryCourseForApp("", "")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetUserTertiaryCourses(t *testing.T) {
	result, err := GetUserTertiaryCourses()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestDeleteUserTertiaryCourse(t *testing.T) {
	obj := domain.UserTertiaryCourse{}
	result, err := DeleteUserTertiaryCourse(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestUpdateUserTertiaryCourse(t *testing.T) {
	obj := domain.UserTertiaryCourse{}
	result, err := UpdateUserTertiaryCourse(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
