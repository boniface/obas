package demographics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/demographics"
	"testing"
)

func TestGetRoles(t *testing.T) {
	value, err := GetRoles()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetRole(t *testing.T) {
	expected := "ADMIN"
	value, err := GetRole("6")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, expected, value.RoleName)
}

func TestCreateRole(t *testing.T) {
	role := domain.Roles{"21", "GUEST"}
	value, err := CreateRole(role)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateRole(t *testing.T) {
	role := domain.Roles{"22", "GUEST"}
	value, err := UpdateRole(role)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteRole(t *testing.T) {
	role := domain.Roles{"21", "GUEST"}
	value, err := DeleteRole(role)
	assert.Nil(t, err)
	assert.True(t, value)
}
