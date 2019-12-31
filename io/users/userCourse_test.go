package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

func TestCreateUserCourse(t *testing.T) {
	obj := domain.UserCourse{"09903", "020289", "0000"}
	value, err := CreateUserCourse(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value)
}
func TestDeleteUserCourse(t *testing.T) {

}
func TestGetUserCourse(t *testing.T) {

}
func TestGetUserCourses(t *testing.T) {

}
func TestUpdateUserCourse(t *testing.T) {
	value, err := GetUserCourses()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value)
}
