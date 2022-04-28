package service

import (
	"atmcase/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_checkPassword_Fail(t *testing.T) {
	pass := "abcd1234"
	idAccount := 123412341234
	check, account := checkPassword(pass, idAccount, model.ListAccount)
	assert.False(t, check)
	assert.Nil(t, account)
}

func TestMain_checkPassword(t *testing.T) {
	pass := "123456"
	idAccount := 123412341234
	check, account := checkPassword(pass, idAccount, model.ListAccount)
	assert.True(t, check)
	assert.Equal(t, account.IdAccount, idAccount)
}

func TestMain_checkIdAccount_FalseInput(t *testing.T) {
	check, id := checkIdAccount("abcd1234", model.ListAccount)
	assert.False(t, check, "False")
	assert.Equal(t, id, 0)
}

func TestMain_checkIdAccount_NotFound(t *testing.T) {
	check, id := checkIdAccount("101012121313", model.ListAccount)
	assert.False(t, check, "False")
	assert.Equal(t, id, 101012121313)
}

func TestMain_checkIdAccount_Ok(t *testing.T) {
	check, id := checkIdAccount("123412341234", model.ListAccount)
	assert.Equal(t, check, true)
	assert.True(t, check)
	assert.IsType(t, id, 1)
	assert.Equal(t, id, 123412341234)
}
