package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/shopspring/decimal"
)

type AccountBank struct {
	IdAccount  int                  `json:"idAccount"`
	PinAccount int                  `json:"pinAccount"`
	Name       string               `json:"name"`
	Balance    decimal.Decimal      `json:"balance"`
	History    []HistoryTransaction `json:"history"`
}

type HistoryTransaction struct {
	Date        string          `json:"date"`
	Transaction string          `json:"transaction"`
	Amount      decimal.Decimal `json:"amount"`
	LastBalance decimal.Decimal `json:"lastBalance"`
}

// fungsi untuk validasi saldo
func (account *AccountBank) ValidasiSaldo(tarik int) bool {
	return account.Balance.GreaterThanOrEqual(decimal.NewFromInt(int64(tarik)))
}

// fungsi untuk menyiapkan data dari file json
func ReadDataAccountJson(jsonName string) []*AccountBank {
	jsonFile, err := os.Open(jsonName)
	if err != nil {
		log.Fatal("gagal baca file json")
	}
	defer jsonFile.Close()

	byteJsonFile, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("gagal convert byte file json")
	}

	var dataAccount []*AccountBank
	err = json.Unmarshal(byteJsonFile, &dataAccount)
	if err != nil {
		log.Fatal("gagal unmarshal file json")
	}
	return dataAccount
}
