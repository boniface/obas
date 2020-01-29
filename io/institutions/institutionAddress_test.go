package institutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/institutions"
	"testing"
)

var obj = domain.InstitutionAddress{"0002", "93939", "20 reibeek stre", "7038"}

func TestCreateInstitutionAddress(t *testing.T) {
	resp, err := CreateInstitutionAddress(obj)
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestDeleteInstitutionAddress(t *testing.T) {
	resp, err := DeleteInstitutionAddress(obj)
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestGetInstitutionAddresses(t *testing.T) {
	resp, err := GetInstitutionAddresses()
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestReadInstitutionAddress(t *testing.T) {
	resp, err := ReadInstitutionAddress("00404", "84847")
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
