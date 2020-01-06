package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

func TestCreateUserTertiarySubject(t *testing.T) {
	obj := domain.UserTertiarySubject{}
	result, err := CreateUserTertiarySubject(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestDeleteUserApplTertiarySubject(t *testing.T) {
	result, err := DeleteUserApplTertiarySubject("", "")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestUpdateUserTertiarySubject(t *testing.T) {
	obj := domain.UserTertiarySubject{}
	result, err := UpdateUserTertiarySubject(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetUserTertiarySubjectGetForApp(t *testing.T) {

}
