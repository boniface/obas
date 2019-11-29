package demographics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/demographics"
	"testing"
)

func TestGetDistricts(t *testing.T) {
	result, err := GetGenders()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetDistrict(t *testing.T) {
	result, err := GetDistrict("0001")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestDeleteDistrict(t *testing.T) {
	obj := domain.District{"0001", "zoneboem"}
	result, err := DeleteDistrict(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestCreateDistrict(t *testing.T) {
	obj := domain.District{"0001", "zoneboem"}
	result, err := CreateDistrict(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestUpdateDistrict(t *testing.T) {
	obj := domain.District{"0001", "zoneboem"}
	result, err := UpdateDistrict(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
