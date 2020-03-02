package location

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/location"
	"strconv"
	"testing"
)

func TestGetLocations(t *testing.T) {
	value, err := GetLocations()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value)
}

func TestGetLocation(t *testing.T) {
	value, err := GetLocation("OOII-CDSIX")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	fmt.Println(" The Results", value.Name)
	assert.NotNil(t, value)
}

func TestCreateSchool(t *testing.T) {
	loc := domain.Location{}
	value, err := CreateLocation(loc)
	assert.Nil(t, err)
	assert.NotNil(t, value)
}

func TestUpdateDocument(t *testing.T) {
	loc := domain.Location{}
	value, err := UpdateLocation(loc)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteDocument(t *testing.T) {
	loc := domain.Location{}
	value, err := DeleteLocation(loc)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestGetLocationsForParent(t *testing.T) {
	value, err := GetLocationsForParent("DACI-MPCGU")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value)
}
func TestCreateLocation(t *testing.T) {
	//var filepath =""
	array_name := [6]string{"A", "B", "C", "D", "E", "F"}
	value, err := GetLocations()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value, "<<<<values")

	f := excelize.NewFile()
	// Create a new sheet.
	//index := f.NewSheet("Sheet2")
	// Set value of a cell.
	for index, value1 := range value {
		indexString := strconv.Itoa(index + 1)
		f.SetCellValue("Sheet1", array_name[0]+indexString, value1.LocationId)

		f.SetCellValue("Sheet1", array_name[1]+indexString, value1.Longitude)

		f.SetCellValue("Sheet1", array_name[2]+indexString, value1.Latitude)

		f.SetCellValue("Sheet1", array_name[3]+indexString, value1.Name)

		f.SetCellValue("Sheet1", array_name[4]+indexString, value1.LocationTypeId)

		f.SetCellValue("Sheet1", array_name[5]+indexString, value1.LocationParentId)
		fmt.Println(value1)

		//nuvo:=array_name[index]+indexString
		if err := f.SaveAs("C:/Users/Nicole Abrahams/go/src/obas/util/files/Location.xlsx"); err != nil {
			fmt.Println(err)
		}
	}
	//f.SetCellValue("Sheet1", "A1", "passionfruit")
	// Set active sheet of the workbook.
	//f.SetActiveSheet(index)
	// Save xlsx file by the given path.

}
