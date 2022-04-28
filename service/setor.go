package service

import (
	"atmcase/model"
	"atmcase/utils"
	"fmt"

	"github.com/shopspring/decimal"
)

const limaPuluhRibu = 50000
const angkaNol = 0

// fungsi untuk proses setor uang
func SetorUang(account *model.AccountBank) {
	fmt.Println("\nSetor Uang :")
	fmt.Println("Setor hanya menerima kelipatan 50000")
	fmt.Printf("\nMasukan jumlah nominal : ")

	nominal, _ := utils.InputScan()
	setor, err := decimal.NewFromString(nominal)
	if err != nil {
		fmt.Println("\n>>>>>> Masukan Salah ! <<<<<<")
		return
	}

	if setor.Mod(decimal.NewFromInt32(limaPuluhRibu)).Equal(decimal.NewFromInt32(angkaNol)) {
		fmt.Println("\nAnda akan melakukan setor uang", nominal)
		fmt.Printf("\n[1] Ya / [0] Tidak : ")
		opsi, _ := utils.InputScan()
		if opsi == "1" {
			result := setor.Add(account.Balance)
			account.Balance = result
			account.History = append(account.History, model.HistoryTransaction{
				Date:        utils.TimeDateNow(),
				Transaction: "Setor",
				Amount:      setor,
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
