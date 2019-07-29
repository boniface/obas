package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/mail"
	"testing"
)

func TestGetSmtpConfig(t *testing.T) {
	expected := "AWS"
	value, err := GetSmtpConfig("12")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.Host, expected)
}

func TestGetSmtpConfigs(t *testing.T) {
	value, err := GetSmtpConfigs()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestCreateSmtpConfig(t *testing.T) {
	smtp := domain.SmtpConfig{"12", "8888", "AWS", "ADMIN", "x1s2s3"}
	value, err := CreateSmtpConfig(smtp)
	assert.Nil(t, err)
	assert.True(t, value)

}

func TestUpdateSmtpConfig(t *testing.T) {
	smtp := domain.SmtpConfig{"15", "8888", "AWS", "ADMIN", "x1s2s3"}
	value, err := UpdateSmtpConfig(smtp)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteSmtpConfig(t *testing.T) {
	smtp := domain.SmtpConfig{"15", "8888", "AWS", "ADMIN", "x1s2s3"}
	value, err := DeleteSmtpConfig(smtp)
	assert.Nil(t, err)
	assert.True(t, value)
}
