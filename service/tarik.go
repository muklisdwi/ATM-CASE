package service

import (
	"atmcase/model"
	"atmcase/utils"
	"fmt"

	"github.com/shopspring/decimal"
)

// fungsi untuk proses tarik uang
func TarikUang(account *model.AccountBank) {
	var isExit bool = false
	var tarik int = 0
	for {
		fmt.Println("\nTarik Uang :")
		fmt.Println("1. 50000")
		fmt.Println("2. 100000")
		fmt.Println("3. 200000")
		fmt.Println("4. 300000")
		fmt.Println("0. Keluar")
		fmt.Printf("\nPilih Menu [1]/[2]/[3]/[4]/[0] : ")
		option, _ := utils.InputScan()
		switch option {
		case "0":
			isExit = true
		case "1":
			tarik = 50000
			isExit = true
		case "2":
			tarik = 100000
			isExit = true
		case "3":
			tarik = 200000
			isExit = true
		case "4":
			tarik = 300000
			isExit = true
		default:
			fmt.Printf("\n>>>>>> Pilihan Salah ! <<<<<<")
			fmt.Println()
		}

		if tarik != 0 {
			if account.ValidasiSaldo(tarik) {
				fmt.Println("\nAnda akan melakukan tarik uang", tarik)
				fmt.Printf("\n[1] Ya / [0] Tidak : ")
				opsi, _ := utils.InputScan()
				if opsi == "1" {
					fmt.Println("\n>>>>> Penarikan diproses <<<<")
					result := account.Balance.Sub(decimal.NewFromInt(int64(tarik)))
					account.Balance = result
					account.History = append(account.History, model.HistoryTransaction{
						Date:        utils.TimeDateNow(),
						Transaction: "Tarik",
						Amount:      decimal.NewFromInt(int64(tarik)),
						Balance:     account.Balance,
					})
				} else {
					fmt.Println("\nTransaksi dibatalkan !")
				}
			} else {
				fmt.Println("\nSaldo tidak cukup !")
			}
		}

		if isExit {
			break
		}
	}
}
