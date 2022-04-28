package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestMain_validasiSaldo(t *testing.T) {
	for _, l := range listAccount {
		check := validasiSaldo(l, 50000)
		assert.True(t, check)
	}
}

func TestMain_findAccount_NotFound(t *testing.T) {
	id := 333344445555
	check, account := findAccount(id)
	assert.False(t, check)
	assert.Nil(t, account)
}

func TestMain_findAccount(t *testing.T) {
	id := 123412341234
	check, account := findAccount(id)
	assert.True(t, check)
	assert.Equal(t, account, listAccount[2])
	assert.Equal(t, account.IdAccount, id)
}

func TestMain_checkPassword_Fail(t *testing.T) {
	pass := "abcd1234"
	idAccount := 123412341234
	check, account := checkPassword(pass, idAccount)
	assert.False(t, check)
	assert.Nil(t, account)
}

func TestMain_checkPassword(t *testing.T) {
	pass := "123456"
	idAccount := 123412341234
	check, account := checkPassword(pass, idAccount)
	assert.True(t, check)
	assert.Equal(t, account.IdAccount, idAccount)
}

func TestMain_checkIdAccount_FalseInput(t *testing.T) {
	check, id := checkIdAccount("abcd1234")
	assert.False(t, check, "False")
	assert.Equal(t, id, 0)
}

func TestMain_checkIdAccount_NotFound(t *testing.T) {
	check, id := checkIdAccount("101012121313")
	assert.False(t, check, "False")
	assert.Equal(t, id, 101012121313)
}

func TestMain_checkIdAccount_Ok(t *testing.T) {
	check, id := checkIdAccount("123412341234")
	assert.Equal(t, check, true)
	assert.True(t, check)
	assert.IsType(t, id, 1)
	assert.Equal(t, id, 123412341234)
}
