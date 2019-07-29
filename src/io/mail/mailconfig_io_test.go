package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/mail"
	"testing"
	"time"
)

func TestGetMailConfigs(t *testing.T) {
	value, err := GetMailConfigs()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetMailConfig(t *testing.T) {
	expected := "SENT"
	value, err := GetMailConfig("UBER")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.State, expected)

}

func TestCreateMailConfig(t *testing.T) {
	config := domain.MailConfig{"36", "PK", "hello", "YAHOO", "212", "SENT", ""}
	value, err := CreateMailConfig(config)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateMailConfig(t *testing.T) {
	config := domain.MailConfig{"37", "PK", "hello", "YAHOO", "212", "SENT", ""}
	value, err := UpdateMailConfig(config)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteMailConfig(t *testing.T) {
	config := domain.MailConfig{"36", "PK", "hello", "YAHOO", "212", "SENT", ""}
	value, err := DeleteMailConfig(config)
	assert.Nil(t, err)
	assert.True(t, value)
}
