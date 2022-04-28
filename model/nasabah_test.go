package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNasabahJson(t *testing.T) {
	var dataAccount []AccountBank
	dataJson := `[{"idAccount":111122223333,"pinAccount":123456,"name":"Akun1","balance":"500000","history":[{"date":"2022-04-27 09:00:14","transaction":"Setor","amount":"500000","lastBalance":"500000"}]}]`
	err := json.Unmarshal([]byte(dataJson), &dataAccount)
	assert.Nil(t, err)
	valueDataAccount, err := json.Marshal(dataAccount)
	if err != nil {
		t.Error("ada error saat marshal")
	}
	assert.Equal(t, dataJson, string(valueDataAccount), "Sama kah ?")
}

// type AccountBank struct {
// 	IdAccount  int                  `json:"idAccount"`
// 	PinAccount int                  `json:"pinAccount"`
// 	Name       string               `json:"name"`
// 	Balance    decimal.Decimal      `json:"balance"`
// 	History    []HistoryTransaction `json:"history"`
// }

// type HistoryTransaction struct {
// 	Date        string          `json:"date"`
// 	Transaction string          `json:"transaction"`
// 	Amount      decimal.Decimal `json:"amount"`
// 	LastBalance decimal.Decimal `json:"lastBalance"`
// }

// {
// 	IdAccount:  111122223333,
// 	PinAccount: 123456,
// 	Name:       "Akun1",
// 	Balance:    decimal.RequireFromString("500000"),
// 	History: []HistoryTransaction{
// 		{
// 			Date:        utils.TimeDateNow(),
// 			Transaction: "Setor",
// 			Amount:      decimal.RequireFromString("500000"),
// 			LastBalance: decimal.RequireFromString("500000"),
// 		},
// 	},
// },
// {
// 	IdAccount:  222233334444,
// 	PinAccount: 123456,
// 	Name:       "Akun2",
// 	Balance:    decimal.RequireFromString("1000000"),
// 	History: []HistoryTransaction{
// 		{
// 			Date:        utils.TimeDateNow(),
// 			Transaction: "Setor",
// 			Amount:      decimal.RequireFromString("1000000"),
// 			LastBalance: decimal.RequireFromString("1000000"),
// 		},
// 	},
// },
// {
// 	IdAccount:  123412341234,
// 	PinAccount: 123456,
// 	Name:       "Akun3",
// 	Balance:    decimal.RequireFromString("1500000"),
// 	History: []HistoryTransaction{
// 		{
// 			Date:        utils.TimeDateNow(),
// 			Transaction: "Setor",
// 			Amount:      decimal.RequireFromString("1500000"),
// 			LastBalance: decimal.RequireFromString("1500000"),
// 		},
// 	},
// },
