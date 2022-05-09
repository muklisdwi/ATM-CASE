package service

import (
	"atmcase/model"
	"atmcase/utils"
	"fmt"

	"github.com/shopspring/decimal"
)

// fungsi untuk proses setor uang
func SetorUang(account *model.AccountBank) {
	fmt.Println("\nSetor Uang :")
	fmt.Println("Setor hanya menerima kelipatan 50000")
	fmt.Printf("\nMasukan jumlah nominal : ")

	nominal, _ := utils.InputScan()
	setorNominal, err := decimal.NewFromString(nominal)
	if err != nil {
		fmt.Println("\n>>>>>> Masukan Salah ! <<<<<<")
		return
	}

	if setorNominal.Mod(decimal.NewFromInt32(utils.LimaPuluhRibu)).Equal(decimal.NewFromInt32(utils.AngkaNol)) {
		fmt.Println("\nAnda akan melakukan setor uang", nominal)
		fmt.Printf("\n[1] Ya / [0] Tidak : ")
		opsi, _ := utils.InputScan()
		if opsi == utils.OpsiYa {
			result := setorNominal.Add(account.Balance)
			account.Balance = result
			account.History = append(account.History, model.HistoryTransaction{
				Date:        utils.TimeDateNow(),
				Transaction: utils.Setor,
				Amount:      setorNominal,
				LastBalance: result,
			})
			fmt.Println("\n>>>>>> Setoran diproses <<<<<")
		} else {
			fmt.Println("\nTransaksi dibatalkan !")
		}
	} else {
		fmt.Println("\n>>>>>> Nominal Salah ! <<<<<<")
		return
	}
}
