package service

import (
	"atmcase/model"
	"atmcase/utils"
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

// proses transfer uang
func TransferUang(accountPengirim *model.AccountBank) {
	fmt.Println("\nTransfer :")
	fmt.Println("Masukan nomor rekening tujuan :")
	strNorek, _ := utils.InputScan()
	intNorek, err := strconv.Atoi(strNorek)
	if err != nil {
		fmt.Printf("\n>>>>>> Masukan Salah ! <<<<<<")
		return
	}
	if check, accountPenerima := findAccount(intNorek); check {
		fmt.Println("\nMasukan nominal transfer : ")
		strNominal, _ := utils.InputScan()
		intNominal, err := strconv.Atoi(strNominal)
		if err != nil {
			fmt.Println("\n>>>>>> Nominal Salah ! <<<<<<")
			return
		}
		if accountPengirim.ValidasiSaldo(intNominal) {
			fmt.Println("\nAnda akan melakukan transfer uang", strNominal, "\nke rekening",
				accountPenerima.IdAccount, "a/n", accountPenerima.Name)
			fmt.Printf("\n[1] Ya / [0] Tidak : ")
			opsi, _ := utils.InputScan()
			if opsi == "1" {
				decimalNominal := decimal.NewFromInt(int64(intNominal)).Abs()
				accountPengirim.Balance = accountPengirim.Balance.Sub(decimalNominal)
				accountPengirim.History = append(accountPengirim.History, model.HistoryTransaction{
					Date:        utils.TimeDateNow(),
					Transaction: "Kirim",
					Amount:      decimalNominal,
					Balance:     accountPengirim.Balance,
				})
				accountPenerima.Balance = accountPenerima.Balance.Add(decimalNominal)
				accountPenerima.History = append(accountPenerima.History, model.HistoryTransaction{
					Date:        utils.TimeDateNow(),
					Transaction: "Terima",
					Amount:      decimalNominal,
					Balance:     accountPenerima.Balance,
				})
				fmt.Println("\n>>>>> Transfer Berhasil ! <<<<")
			} else {
				fmt.Println("\nTransaksi dibatalkan !")
			}
		} else {
			fmt.Println("\n>>>>> Saldo tidak cukup ! <<<<")
		}
	} else {
		fmt.Println("\n>>> Akun tidak ditemukan ! <<<")
	}
}

// fungsi untuk cari rekening tujuan
func findAccount(id int) (bool, *model.AccountBank) {
	var check bool = false
	var account *model.AccountBank
	for i, l := range model.ListAccount {
		if id == l.IdAccount {
			check = true
			account = model.ListAccount[i]
		}
	}
	return check, account
}
