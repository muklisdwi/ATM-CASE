package model

import "github.com/shopspring/decimal"

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
