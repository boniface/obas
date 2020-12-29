package applications

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/application"
	"testing"
	"time"
)

func TestGetAllStatusesForApplication(t *testing.T) {
	value, err := GetAllStatusesForApplication("MMMO-4HJVW")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value)
}

func TestGetApplicationStatus(t *testing.T) {
	//expected := "Pending"
	value, err := GetApplicationStatus("MMMO-4HJVW")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value.ApplicationId)
}
func TestGetAllStatusesForApplication2(t *testing.T) {

}

func TestCreateApplicationStatus(t *testing.T) {
	//appType := domain.ApplicationStatus{"212", "Successful", "Final decision"}
	//appType := application.ApplicationStatus{"CCMO-0TAGM", "Completed", "Done", "", time.Now()}
	appType := domain.ApplicationStatus{"CCMO-0TAGM", "Completed", "Done", "", time.Now()}
	value, err := CreateApplicationStatus(appType)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}
