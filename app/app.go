package app

import (
	"atmcase/model"
	"atmcase/service"
	"atmcase/utils"
	"fmt"
	"os"
)

// fungsi untuk menjalankan aplikasi
func StartApp() {
	fmt.Println("\n>>>>>> Selamat Datang ! <<<<<<")
	for {
		OptionLogin()
	}
}

// menampilkan menu login
func OptionLogin() {
	fmt.Println("\nSilahakan Login !")
	fmt.Println("\n1. Login")
	fmt.Println("0. Keluar")
	fmt.Printf("\nPilih Menu [1]/[0] : ")
	option, _ := utils.InputScan()
	switch option {
	case "0":
		fmt.Println("\n>>>>>> Terima Kasih ! <<<<<<")
		fmt.Println()
		os.Exit(0)
	case "1":
		isAccount := service.LoginProccess(model.ListAccount)
		if isAccount != nil {
			OptionAccount(isAccount)
		} else {
			fmt.Println("Login Gagal !")
		}
	default:
		fmt.Printf("\n>>>>>> Pilihan Salah ! <<<<<<")
		fmt.Println()
	}
}

// menampilkan menu akun setelah login berhasil
func OptionAccount(account *model.AccountBank) {
	var isLogout bool = false
	for {
		fmt.Println("\nSelamat Datang,", account.Name)
		fmt.Println("1. Tarik Uang")
		fmt.Println("2. Setor Uang")
		fmt.Println("3. Transfer")
		fmt.Println("4. Lihat Riwayat Transaksi")
		fmt.Println("0. Keluar")
		fmt.Printf("\nPilih Menu [1]/[2]/[3]/[4]/[0] : ")
		option, _ := utils.InputScan()
		switch option {
		case "0":
			fmt.Printf("\n")
			fmt.Println(account.Name, "telah keluar")
			fmt.Println("Terima Kasih")
			isLogout = true
		case "1":
			service.TarikUang(account)
		case "2":
			service.SetorUang(account)
		case "3":
			service.TransferUang(account)
		case "4":
			service.RiwayatTransaksi(account)
		default:
			fmt.Println("\n>>>>>> Pilihan Salah ! <<<<<<")
		}

		if isLogout {
			break
		}
	}
}
