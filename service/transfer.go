package service

import (
	"atmcase/model"
	"atmcase/utils"
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

// proses transfer uang
func TransferUang(accountPengirim *model.AccountBank, listAccount []*model.AccountBank) {
	fmt.Println("\nTransfer :")
	fmt.Println("Masukan nomor rekening tujuan :")
	strNorek, _ := utils.InputScan()

	intNorek, err := strconv.Atoi(strNorek)
	if err != nil {
		fmt.Printf("\n>>>>>> Masukan Salah ! <<<<<<")
		return
	}

	check, accountPenerima := findAccount(intNorek, listAccount)
	if !check {
		fmt.Println("\n>>> Akun tidak ditemukan ! <<<")
		return
	}

	fmt.Println("\nMasukan nominal transfer : ")
	strNominal, _ := utils.InputScan()

	intNominal, err := strconv.Atoi(strNominal)
	if err != nil {
		fmt.Println("\n>>>>>> Nominal Salah ! <<<<<<")
		return
	}

	isSaldoValid := accountPengirim.ValidasiSaldo(intNominal)
	if !isSaldoValid {
		fmt.Println("\n>>>>> Saldo tidak cukup ! <<<<")
		return
	}

	fmt.Println("\nAnda akan melakukan transfer uang", strNominal, "\nke rekening",
		accountPenerima.IdAccount, "a/n", accountPenerima.Name)

	fmt.Printf("\n[1] Ya / [0] Tidak : ")
	opsi, _ := utils.InputScan()
	if opsi == utils.OpsiYa {
		decimalNominal := decimal.NewFromInt(int64(intNominal)).Abs()
		accountPengirim.Balance = accountPengirim.Balance.Sub(decimalNominal)
		accountPengirim.History = append(accountPengirim.History, model.HistoryTransaction{
			Date:        utils.TimeDateNow(),
			Transaction: utils.Kirim,
			Amount:      decimalNominal,
			LastBalance: accountPengirim.Balance,
		})
		accountPenerima.Balance = accountPenerima.Balance.Add(decimalNominal)
		accountPenerima.History = append(accountPenerima.History, model.HistoryTransaction{
			Date:        utils.TimeDateNow(),
			Transaction: utils.Terima,
			Amount:      decimalNominal,
			LastBalance: accountPenerima.Balance,
		})
		fmt.Println("\n>>>>> Transfer Berhasil ! <<<<")
	} else {
		fmt.Println("\nTransaksi dibatalkan !")
	}
}

// fungsi untuk cari rekening tujuan
func findAccount(id int, listAccount []*model.AccountBank) (bool, *model.AccountBank) {
	var check bool = false
	var account *model.AccountBank
	for i, l := range listAccount {
		if id == l.IdAccount {
			check = true
			account = listAccount[i]
			break
		}
	}
	return check, account
}
