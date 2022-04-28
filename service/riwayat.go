package service

import (
	"atmcase/model"
	"atmcase/utils"
	"fmt"
)

// lihat daftar riwayat transaksi
func RiwayatTransaksi(account *model.AccountBank) {
	fmt.Println("\nDaftar Riwayat Transaksi :")
	for _, l := range account.History {
		fmt.Printf("%v | %s | %v | %v \n", l.Date, l.Transaction, l.Amount, l.Balance)
	}
	fmt.Printf("\n(Tekan Enter)")
	utils.InputScan()
}
