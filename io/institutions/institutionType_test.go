package institutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/institutions"
	"testing"
)

func TestCreateInstitutionType(t *testing.T) {
	obj := domain.InstitutionTypes{"001", "UNIVERSITY", ""}
	resp, err := CreateInstitutionType(obj)
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestDeleteInstitutionType(t *testing.T) {
	obj := domain.InstitutionTypes{"001", "UNIVERSITY", ""}
	resp, err := DeleteInstitutionType(obj)
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestGetInstitutionType(t *testing.T) {
	resp, err := GetInstitutionType("ooo1")
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestGetInstitutionTypes(t *testing.T) {
	resp, err := GetInstitutionTypes()
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestUpdateInstitutionType(t *testing.T) {
	obj := domain.InstitutionTypes{"001", "UNIVERSITY", ""}
	resp, err := UpdateInstitutionType(obj)
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
