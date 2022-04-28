package service

import (
	"atmcase/model"
	"atmcase/utils"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

var listAccount = []*model.AccountBank{
	{
		IdAccount:  123412341234,
		PinAccount: 123456,
		Name:       "Akun3",
		Balance:    decimal.NewFromInt32(1500000),
		History: []model.HistoryTransaction{
			{
				Date:        utils.TimeDateNow(),
				Transaction: "Setor",
				Amount:      decimal.NewFromInt32(1500000),
				LastBalance: decimal.NewFromInt32(1500000),
			},
		},
	},
}

func TestMain_checkPassword_Fail(t *testing.T) {
	pass := "abcd1234"
	idAccount := 123412341234
	check, account := checkPassword(pass, idAccount, listAccount)
	assert.False(t, check)
	assert.Nil(t, account)
}

func TestMain_checkPassword(t *testing.T) {
	pass := "123456"
	idAccount := 123412341234
	check, account := checkPassword(pass, idAccount, listAccount)
	assert.True(t, check)
	assert.Equal(t, account.IdAccount, idAccount)
}

func TestMain_checkIdAccount_FalseInput(t *testing.T) {
	check, id := checkIdAccount("abcd1234", listAccount)
	assert.False(t, check, "False")
	assert.Equal(t, id, 0)
}

func TestMain_checkIdAccount_NotFound(t *testing.T) {
	check, id := checkIdAccount("101012121313", listAccount)
	assert.False(t, check, "False")
	assert.Equal(t, id, 101012121313)
}

func TestMain_checkIdAccount_Ok(t *testing.T) {
	check, id := checkIdAccount("123412341234", listAccount)
	assert.Equal(t, check, true)
	assert.True(t, check)
	assert.IsType(t, id, 1)
	assert.Equal(t, id, 123412341234)
}
