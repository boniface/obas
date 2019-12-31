package institutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/institutions"
	"testing"
)

var entity = domain.InstitutionTypes{Id: "1991", Name: "Christian M", Description: ""}

func TestCreateInstitutionType(t *testing.T) {
	obj := domain.InstitutionTypes{"001", "College", "High certificate"}
	resp, err := CreateInstitutionType(obj)
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestDeleteInstitutionType(t *testing.T) {
	obj := domain.InstitutionTypes{"EELO-LMEIR", "College", "High certificate"}
	resp, err := DeleteInstitutionType(obj)
	assert.True(t, resp)
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
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, len(resp) > 0)
}
func TestUpdateInstitutionType(t *testing.T) {
	obj := domain.InstitutionTypes{"EEEE-IKMRY", "College", "this is a colleges for further education"}
	resp, err := UpdateInstitutionType(obj)
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}

/*func TestUpdateContactType(t *testing.T) {
var expected = "Christian Muamba"
var updated = domain.ContactType{ContactTypeId: "1991", Name: "Christian Muamba"}
result, err := UpdateContactType(updated)
assert.Nil(t, err)
assert.True(t, result)
value, err := GetContactType(entity.ContactTypeId)
assert.Equal(t, expected, value.Name)*/
