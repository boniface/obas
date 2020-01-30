package demographics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/demographics"
	"testing"
)

var token = "eyJraWQiOiJURVNUX1BIUkFTRSIsImFsZyI6IkVTMjU2In0.eyJpc3MiOiJIQVNIQ09ERS5aTSIsImF1ZCI6IlNJVEVVU0VSUyIsImV4cCI6MTU4MDIxNTU2OCwianRpIjoiX284WGVZX1h3WGNVU2pTWjZMTGZkdyIsImlhdCI6MTU4MDEyOTE2OCwibmJmIjoxNTgwMTI5MDQ4LCJzdWIiOiJTaXRlIEFjY2VzcyIsImVtYWlsIjoiZXNwb2lyZGl0ZWtlbWVuYUBnbWFpbC5jb20iLCJyb2xlIjoiQUFJSS05Q1pEViJ9.2FCbuRUZbFygGDD7KoGiEpYlIWhgz6b2IZ8_n1x3m3NObL47eLn6uFbkCxy26UPkA-RH3ylcDJeHPBltd3w8MA"

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
	role := domain.Role{"", "GUEST"}
	value, err := CreateRole(role)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateRole(t *testing.T) {
	role := domain.Role{"1", "testes"}
	value, err := UpdateRole(role, token)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteRole(t *testing.T) {
	role := domain.Role{"ESTT-15DXY", "testes"}
	value, err := DeleteRole(role)
	assert.Nil(t, err)
	assert.True(t, value)
}
