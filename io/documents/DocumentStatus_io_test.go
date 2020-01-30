package documents

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/documents"
	"testing"
	"time"
)

func TestCreateDocumentStatus(t *testing.T) {
	obj := domain.DocumentStatus{"03030", "848484", "4975834204924", "nothing", time.Now()}
	result, err := CreateDocumentStatus(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetdocumentStatues(t *testing.T) {
	result, err := GetdocumentStatues("03030")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetDocumentStatus(t *testing.T) {
	result, err := GetDocumentStatus("03030")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
