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
	var tarikNominal int = 0
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
		case utils.CaseNol:
			isExit = true
		case utils.CaseSatu:
			tarikNominal = utils.LimaPuluhRibu
			isExit = true
		case utils.CaseDua:
			tarikNominal = utils.SeratusRibu
			isExit = true
		case utils.CaseTiga:
			tarikNominal = utils.DuaRatusRibu
			isExit = true
		case utils.CaseEmpat:
			tarikNominal = utils.TigaRatusRibu
			isExit = true
		default:
			fmt.Printf("\n>>>>>> Pilihan Salah ! <<<<<<")
			fmt.Println()
			isExit = true
		}

		if tarikNominal == 0 && isExit {
			break
		}

		if !account.ValidasiSaldo(tarikNominal) {
			fmt.Println("\nSaldo tidak cukup !")
			continue
		}

		fmt.Println("\nAnda akan melakukan tarik uang", tarikNominal)
		fmt.Printf("\n[1] Ya / [0] Tidak : ")
		opsi, _ := utils.InputScan()
		if opsi == utils.OpsiYa {
			fmt.Println("\n>>>>> Penarikan diproses <<<<")
			result := account.Balance.Sub(decimal.NewFromInt(int64(tarikNominal)))
			account.Balance = result
			account.History = append(account.History, model.HistoryTransaction{
				Date:        utils.TimeDateNow(),
				Transaction: utils.Tarik,
				Amount:      decimal.NewFromInt(int64(tarikNominal)),
				LastBalance: account.Balance,
			})
		} else {
			fmt.Println("\nTransaksi dibatalkan !")
		}

		if isExit {
			break
		}
	}
}
