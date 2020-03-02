package documents

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/documents"
	"testing"
)

var token = "eyJraWQiOiJvYmFzYXBpX29uXzE1NS4yMzguMzIuMjE5IiwiYWxnIjoiRVMyNTYifQ.eyJpc3MiOiJIQVNIQ09ERS5aTSIsImF1ZCI6IlNJVEVVU0VSUyIsImV4cCI6MTU4MDU3MDU0NiwianRpIjoiMWlqaXFnWVJXbWxFdlgtd09hSWozQSIsImlhdCI6MTU4MDQ4NDE0NiwibmJmIjoxNTgwNDg0MDI2LCJzdWIiOiJTaXRlIEFjY2VzcyIsImVtYWlsIjoiMjE2MDkzODA1QG15Y3B1dC5hYy56YSIsInJvbGUiOiJFTlNULTEyOVNVIn0.4O0u91m6DoDGLavxUO6IPA2muVzTqFWpxoLwMZyiCU3ei6ttVbtiGW_fjdX9J1pUOdXSSD9WhfCoER1Zo9gPAw"

func TestGetDocumentTypes(t *testing.T) {
	value, err := GetDocumentTypes()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetDocumentsType(t *testing.T) {
	expected := "COURSES"
	value, err := GetDocumentType("LACP-5CEXS")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value, expected)
}
func TestGetDocumentTypes2(t *testing.T) {

}

func TestCreateDocumentTypes(t *testing.T) {
	//docType := domain.DocumentType{"2", "COURSES"}
	docType := domain.DocumentType{"1", "Acceptation "}
	value, err := CreateDocumentType(docType, token)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateDocumentTypes(t *testing.T) {
	docType := domain.DocumentType{"2", "COURSES"}
	value, err := UpdateDocument(docType)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteDocumentTypes(t *testing.T) {
	docType := domain.DocumentType{"2", "COURSES"}
	value, err := DeleteDocumentType(docType)
	assert.Nil(t, err)
	assert.True(t, value)
}
