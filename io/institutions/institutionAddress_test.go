package institutions

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
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

/****This method reads all the institutions Address data from the file and send them the api***/
func TestCreateInstitutionAddress2(t *testing.T) {
	institutionsAddress, err := excelize.OpenFile("C:/Users/Nicole Abrahams/go/src/obas/util/files/institution_address.xlsx")
	var newInstitution domain.InstitutionAddress
	if err != nil {
		fmt.Println(err)
		return
	}
	cellVal, err := institutionsAddress.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	/***Looping through all the rows that contains Data***/
	for _, value := range cellVal {
		/***reading the first value in the first row***/
		newInstitution = domain.InstitutionAddress{value[0], value[1], value[2], value[3]}
		/***Now sending the object to the api***/
		CreateInstitutionAddress(newInstitution)
		/**Now clearing the object**/
		newInstitution = domain.InstitutionAddress{}
	}
}
