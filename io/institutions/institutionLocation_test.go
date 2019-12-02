package institutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/institutions"
	"testing"
)

func TestCreateInstitutionLocation(t *testing.T) {
	obj := domain.InstitutionLocation{"0001", "0002", "-23231", "3234324"}
	resp, err := CreateInstitutionLocation(obj)
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestDeleteInstitutionLocation(t *testing.T) {
	obj := domain.InstitutionLocation{"0001", "0002", "-23231", "3234324"}
	resp, err := DeleteInstitutionLocation(obj)
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestReadInstitutionLocation(t *testing.T) {
	resp, err := ReadInstitutionLocation("0001")
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestUpdateInstitutionLocation(t *testing.T) {
	obj := domain.InstitutionLocation{"0001", "0002", "-23231", "3234324"}
	resp, err := UpdateInstitutionLocation(obj)
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestReadInstitutionLocations(t *testing.T) {
	resp, err := ReadInstitutionLocations()
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
