package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

func TestGetUserDocuments(t *testing.T) {
	result, err := GetUserDocuments("espoirditekemena@gmail.com")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, len(result) > 0)
}

func TestGetUserDocument(t *testing.T) {
	expected := "45"
	result, err := GetUserDocument("12", "")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, expected, result.DocumentId)

}

func TestCreateUserDocument(t *testing.T) {
	userDoc := domain.UserDocument{"585", "532"}
	result, err := CreateUserDocument(userDoc, "")
	assert.Nil(t, err)
	assert.True(t, result)

}

func TestUpdateUserDocument(t *testing.T) {
	userDoc := domain.UserDocument{"12", "45"}
	result, err := UpdateUserDocument(userDoc, "")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)
}

func TestDeleteUserDocument(t *testing.T) {
	userDoc := domain.UserDocument{"585", "532"}
	result, err := DeleteUserDocument(userDoc)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)

}
