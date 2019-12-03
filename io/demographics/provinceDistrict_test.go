package demographics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/demographics"
	"testing"
)

func TestGetProvinceDistrict(t *testing.T) {
	result, err := GetDistrictsInProvince("0001")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetProvinceDistricts(t *testing.T) {
	result, err := GetProvinceDistricts()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestUpdateProvinceDistrict(t *testing.T) {
	obj := domain.ProvinceDistrict{"001", "0001"}
	result, err := UpdateProvinceDistrict(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestDeleteProvinceDistrict(t *testing.T) {
	obj := domain.ProvinceDistrict{"001", "0001"}
	result, err := DeleteProvinceDistrict(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestCreateProvinceDistrict(t *testing.T) {
	obj := domain.ProvinceDistrict{"001", "0001"}
	result, err := CreateProvinceDistrict(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
