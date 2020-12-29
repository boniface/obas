package institutions

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/institutions"
	"testing"
)

func TestCreateInstitutionCourse(t *testing.T) {
	obj := domain.InstitutionCourse{"00012", "0040"}
	resp, err := CreateInstitutionCourse(obj)
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestDeleteInstitutionCourse(t *testing.T) {
	obj := domain.InstitutionCourse{"00012", "0040"}
	resp, err := DeleteInstitutionCourse(obj)
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}

func TestGetInstitutionCourses(t *testing.T) {
	resp, err := GetInstitutionCourses("")
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestReadInstitutionCourse(t *testing.T) {
	resp, err := ReadInstitutionCourse("00012", "0040")
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}

/****This method reads all the institutions Course data from the file and send them the api***/

func TestCreateInstitutionCourse22(t *testing.T) {
	institutionsCourse, err := excelize.OpenFile("C:/Users/Nicole Abrahams/go/src/obas/util/files/institution_course.xlsx")
	var newInstitution domain.InstitutionCourse
	if err != nil {
		fmt.Println(err)
		return
	}
	cellVal, err := institutionsCourse.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	/***Looping through all the rows that contains Data***/
	for _, value := range cellVal {
		/***reading the first value in the first row***/
		newInstitution = domain.InstitutionCourse{value[0], value[1]}
		/***Now sending the object to the api***/
		CreateInstitutionCourse(newInstitution)
		/**Now clearing the object**/
		newInstitution = domain.InstitutionCourse{}
	}
}
