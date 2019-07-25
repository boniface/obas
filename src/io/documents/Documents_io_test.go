package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/documents"
	"testing"
	"time"
)

func TestDocuments(t *testing.T) {
	value, err := GetDocuments()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetDocument(t *testing.T) {
	expected := "Matric Result"
	value, err := GetDocument("56")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.Description, expected)
}

func TestCreateDocument(t *testing.T) {
	doc := domain.Documents{"FG", "25", "FR", "MATRIC", "DS", "QA", "2019", "NONE"}
	value, err := CreateDocument(doc)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateDocument(t *testing.T) {
	doc := domain.Documents{"FG", "25", "FR", "MATRIC", "DS", "QA", time.Time{2019, 03, 1}, "NONE"}
	value, err := UpdateDocument(doc)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteDocument(t *testing.T) {
	doc := domain.Documents{"FG", "25", "FR", "MATRIC", "DS", "QA", 25 / 03 / 2019, "NONE"}
	value, err := DeleteDocument(doc)
	assert.Nil(t, err)
	assert.True(t, value)
}
