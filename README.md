## Test Cases Samples for Obas
```golang
import (
 	"github.com/stretchr/testify/assert"
 	"testing"
 )
 
 
 func TestGetZones(t *testing.T) {
 	result, err := GetZones()
 	assert.Nil(t, err)
 
 	assert.True(t, len(result) > 0)
 }
 
 func TestGetZone(t *testing.T) {
 	expected := "TANZANIA"
 	result, err := GetZone("TZ")
 	assert.Nil(t, err)
 	assert.Equal(t, expected, result.Name)
 }
 
 func TestGetDisabledZones(t *testing.T) {
 	result, err := GetDisabledZones()
 	assert.Nil(t, err)
 
 	assert.True(t, len(result) == 0)
 }
 func TestCreateZone(t *testing.T) {
 	result, err := CreateZone("SITE")
 	assert.Nil(t, err)
 	assert.True(t, result)
 
 }
 func TestDeleteZone(t *testing.T) {
 	result, err := DeleteZone("SITe")
 	assert.Nil(t, err)
 	assert.True(t, result)
 
 }
 func TestUpdateZone(t *testing.T) {
 	result, err := UpdateZone("LT")
 	assert.Nil(t, err)
 	assert.True(t, result)
 
 }
 func TestGetActiveZones(t *testing.T) {
 	result, err := GetActiveZones()
 	assert.Nil(t, err)
 	assert.True(t, len(result) > 0)
 }`
 ```
