package documents

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/documents"
	"time"

	//domain3 "obas/domain/documents"
	"testing"
)

var doc = domain.Document{"test@test.go", "25", "FR", "MATRIC", "DS", time.Now(), "", "NONE"}

//var token=""
func TestDocuments(t *testing.T) {
	value, err := GetDocuments()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetDocument(t *testing.T) {
	expected := doc
	value, err := GetDocument("403859ad7e5cfba0a8818c1bcb60c27f")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value, expected)
}

func TestCreateDocument(t *testing.T) {
	value, err := CreateDocument(doc, "")
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateDocument(t *testing.T) {
	var expected = "MATRIC"
	//var doc = domain.Documents{"FG", "27", "FR", "MATRIC", "DS", "QA", "", "NONE"}
	result, err := UpdateDocument(doc)
	assert.Nil(t, err)
	assert.True(t, result)
	value, err := GetDocument(doc.DocumentId)
	assert.Equal(t, expected, value.Description)
}

func TestDeleteDocument(t *testing.T) {
	value, err := DeleteDocument(doc)
	assert.Nil(t, err)
	assert.True(t, value)
}
