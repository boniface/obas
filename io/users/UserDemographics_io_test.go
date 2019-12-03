package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

func TestGetUserDemographics(t *testing.T) {
	result, err := GetUserDemographics()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, len(result) > 0)
}

func TestGetUserDemographic(t *testing.T) {
	expected := "genderTest"
	result, err := GetUserDemographic("1")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, expected, result.GenderId)

}

func TestCreateUserDemographics(t *testing.T) {
	userDemo := domain.UserDemographics{"516", "215", "826", "848484"}
	result, err := CreateUserDemographics(userDemo)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)

}

//func TestUpdateUserDemographics(t *testing.T) {
//	userDemo := domain.UserDemographics{"516", "215", "826","848484"}
//	result, err := UpdateUserDemographics(userDemo,"")
//	assert.Nil(t, err)
//	assert.True(t, result)
//}

func TestDeleteUserDemographics(t *testing.T) {
	userDemo := domain.UserDemographics{"516", "215", "826", "848484"}
	result, err := DeleteUserDemographics(userDemo)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)

}
