package demographics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/demographics"
	"testing"
)

func TestGetRaces(t *testing.T) {
	value, err := GetRaces()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetRace(t *testing.T) {
	expected := "ASIAN"
	value, err := GetRace("8")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.RaceName, expected)
}

func TestCreateRace(t *testing.T) {
	race := domain.Race{"50", "RED"}
	value, err := CreateRace(race)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateRace(t *testing.T) {
	race := domain.Race{"50", "RED"}
	value, err := UpdateRace(race)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value)
}

func TestDeleteRace(t *testing.T) {
	race := domain.Race{"50", "RED"}
	value, err := DeleteRace(race)
	assert.Nil(t, err)
	assert.NotNil(t, value)
}
