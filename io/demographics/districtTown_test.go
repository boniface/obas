package demographics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/demographics"
	"testing"
)

func TestCreateDistrictTown(t *testing.T) {
	obj := domain.DistrictTown{"000", "0001"}
	result, err := CreateDistrictTown(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestDeleteDistrictTown(t *testing.T) {
	obj := domain.DistrictTown{"000", "0001"}
	result, err := DeleteDistrictTown(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetDistrictTowns(t *testing.T) {
	result, err := GetDistrictTowns()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetDistrictTown(t *testing.T) {
	result, err := GetDistrictTown("0001")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestUpdateDistrictTown(t *testing.T) {
	obj := domain.DistrictTown{"000", "0001"}
	result, err := UpdateDistrictTown(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
