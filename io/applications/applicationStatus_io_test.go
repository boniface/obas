package applications

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/application"
	"testing"
	"time"
)

func TestGetAllStatusesForApplication(t *testing.T) {

}

func TestGetApplicationStatus(t *testing.T) {
	expected := "Pending"
	value, err := GetApplicationStatus("212")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.ApplicationId, expected)
}

func TestCreateApplicationStatus(t *testing.T) {
	//appType := domain.ApplicationStatus{"212", "Successful", "Final decision"}
	appType := domain.ApplicationStatus{"CCMO-0TAGM", "Completed", "Done", "", time.Now()}
	value, err := CreateApplicationStatus(appType)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}
