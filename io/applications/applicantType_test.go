package applications

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/application"
	"testing"
)

func TestCreateApplicantType(t *testing.T) {
	obj := domain.ApplicantType{"EEGG-4CNTX", "College applicant", "A student currently doing his high certificate"}
	valeu, err := CreateApplicantType(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", valeu)
	assert.NotNil(t, valeu)
}
func TestDeleteApplicantType(t *testing.T) {
	obj := domain.ApplicantType{"EILL-3OUVL", "COLLEGE", "with high certificate"}
	valeu, err := DeleteApplicantType(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", valeu)
	assert.NotNil(t, valeu)
}
func TestGetApplicantType(t *testing.T) {
	valeu, err := GetApplicantType("")
	assert.Nil(t, err)
	fmt.Println(" The Results", valeu)
	assert.NotNil(t, valeu)
}
func TestGetApplicantTypes(t *testing.T) {
	valeu, err := GetApplicantTypes()
	assert.Nil(t, err)
	fmt.Println(" The Results", valeu)
	assert.NotNil(t, valeu)
}
