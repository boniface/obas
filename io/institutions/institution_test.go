package institutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/institutions"
	"testing"
)

// {UEVY-IRFGJ UEIN-BXYBY Cape Peninsula University Of Technology}

func TestCreateInstitution(t *testing.T) {
	obj := domain.Institution{"UEVY-IRFGJ", "VVVY-CJKQW", "Cape Peninsula University Of Technology"}
	resp, err := CreateInstitution(obj)
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestGetInstitutions(t *testing.T) {
	resp, err := GetInstitutions()
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, len(resp) > 0)
}
func TestDeleteInstitution(t *testing.T) {
	obj := domain.Institution{"CLNT-BYYLZ", "UEIN-BXYBY", "Cape Peninsula University Of Technology"}
	resp, err := DeleteInstitution(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
func TestGetInstitution(t *testing.T) {
	resp, err := GetInstitution("AAET-08ISZ")
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
func TestGetInstitutionsByType(t *testing.T) {
	resp, err := GetInstitutionsByType("ERSS-4EBHJ")
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
func TestUpdateInstitution(t *testing.T) {
	obj := domain.Institution{"CILN-FSBER", "VVVY-CJKQW", "STELLENBOSCH UNIVERSITY"}
	resp, err := UpdateInstitution(obj, "")
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
