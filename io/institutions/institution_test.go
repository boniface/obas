package institutions

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
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

/****This method reads all the institutions data from the file and send them the api***/
func TestCreateInstitution2(t *testing.T) {
	institutions, err := excelize.OpenFile("C:/Users/Nicole Abrahams/go/src/obas/util/files/institution.xlsx")
	var newInstitution domain.Institution
	if err != nil {
		fmt.Println(err)
		return
	}
	cellVal, err := institutions.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	/***Looping through all the rows that contains Data***/
	for _, value := range cellVal {
		/***reading the first value in the first row***/
		newInstitution = domain.Institution{value[0], value[1], value[2]}
		/***Now sending the object to the api***/
		CreateInstitution(newInstitution)
		/**Now clearing the object**/
		newInstitution = domain.Institution{}
	}
}
