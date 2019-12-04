package applications

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/application"
	"testing"
)

func TestGetApplicationStatuses(t *testing.T) {
	value, err := GetApplicationStatuses()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, len(value) > 0)

}

func TestGetApplicationStatus(t *testing.T) {
	expected := "Pending"
	value, err := GetApplicationStatus("212")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.Id, expected)
}

func TestCreateApplicationStatus(t *testing.T) {
	//appType := domain.ApplicationStatus{"212", "Successful", "Final decision"}
	appType := domain.ApplicationStatus{"CCMO-0TAGM", "Completed", "Done with all the requirements of application process"}
	value, err := CreateApplicationStatus(appType)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}
func TestUpdateApplicationStatus(t *testing.T) {
	appType := domain.ApplicationStatus{"CCMO-0TAGM", "Completed", "Done with all the requirements of application process"}
	value, err := UpdateApplicationStatus(appType)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value)
}
func TestDeleteApplicationStatus(t *testing.T) {
	appType := domain.ApplicationStatus{"ICNP-47IRU", "Incompleted", "New application"}
	value, err := DeleteApplicationStatus(appType)
	assert.Nil(t, err)
	assert.True(t, value)
}
