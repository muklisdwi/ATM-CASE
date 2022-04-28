package model

import (
	"atmcase/utils"

	"github.com/shopspring/decimal"
)

type AccountBank struct {
	IdAccount  int
	PinAccount int
	Name       string
	Balance    decimal.Decimal
	History    []HistoryTransaction
}

type HistoryTransaction struct {
	Date        string
	Transaction string
	Amount      decimal.Decimal
	Balance     decimal.Decimal
}

// fungsi untuk validasi saldo
func (account *AccountBank) ValidasiSaldo(tarik int) bool {
	return account.Balance.GreaterThanOrEqual(decimal.NewFromInt(int64(tarik)))
}

var ListAccount = []*AccountBank{
	{
		IdAccount:  111122223333,
		PinAccount: 123456,
		Name:       "Akun1",
		Balance:    decimal.RequireFromString("500000"),
		History: []HistoryTransaction{
			{
				Date:        utils.TimeDateNow(),
				Transaction: "Setor",
				Amount:      decimal.RequireFromString("500000"),
				Balance:     decimal.RequireFromString("500000"),
			},
		},
	},
	{
		IdAccount:  222233334444,
		PinAccount: 123456,
		Name:       "Akun2",
		Balance:    decimal.RequireFromString("1000000"),
		History: []HistoryTransaction{
			{
				Date:        utils.TimeDateNow(),
				Transaction: "Setor",
				Amount:      decimal.RequireFromString("1000000"),
				Balance:     decimal.RequireFromString("1000000"),
			},
		},
	},
	{
		IdAccount:  123412341234,
		PinAccount: 123456,
		Name:       "Akun3",
		Balance:    decimal.RequireFromString("1500000"),
		History: []HistoryTransaction{
			{
				Date:        utils.TimeDateNow(),
				Transaction: "Setor",
				Amount:      decimal.RequireFromString("1500000"),
				Balance:     decimal.RequireFromString("1500000"),
			},
		},
	},
}
