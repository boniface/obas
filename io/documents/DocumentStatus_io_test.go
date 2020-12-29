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
	result, err := GetdocumentStatues("403859ad7e5cfba0a8818c1bcb60c27f")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetDocumentStatus(t *testing.T) {
	result, err := GetDocumentStatus("403859ad7e5cfba0a8818c1bcb60c27f")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
