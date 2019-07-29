package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/mail"
	"testing"
)

func TestGetMailApis(t *testing.T) {
	value, err := GetMailApis()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetMailApi(t *testing.T) {
	expected := "M8ZH"
	value, err := GetMailApi("147")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.Sender, expected)
}

func TestCreateMailApi(t *testing.T) {
	mail := domain.MailApi{"152", "MK5Y", "hENRY"}
	value, err := CreateMailApi(mail)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateMailApi(t *testing.T) {
	mail := domain.MailApi{"153", "MK5Y", "hENRY"}
	value, err := UpdateMailApi(mail)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestDeleteMailApi(t *testing.T) {
	mail := domain.MailApi{"152", "MK5Y", "hENRY"}
	value, err := DeleteMailApi(mail)
	assert.Nil(t, err)
	assert.True(t, value)
}
