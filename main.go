package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

func main() {
	StartApp()
}

func StartApp() {
	fmt.Println("\n>>>>>> Selamat Datang ! <<<<<<")
	for {
		OptionLogin()
		// fmt.Println(LoginProccess())
		// CheckError(nil)
	}
}

type AccountBank struct {
	IdAccount  int
	PinAccount int
	Name       string
	Balance    decimal.Decimal
	History    []HistoryTransaction
}

type HistoryTransaction struct {
	Date        time.Time
	Transaction string
	Amount      decimal.Decimal
	Balance     decimal.Decimal
}

var listAccount = []AccountBank{
	AccountBank{
		IdAccount:  111122223333,
		PinAccount: 123456,
		Name:       "Akun1",
		Balance:    decimal.RequireFromString("500000"),
		History: []HistoryTransaction{
			HistoryTransaction{
				Date:        time.Date(2022, 04, 26, 8, 00, 00, 00, time.Local),
				Transaction: "Setor",
				Amount:      decimal.RequireFromString("500000"),
				Balance:     decimal.RequireFromString("500000"),
			},
		},
	},
	AccountBank{
		IdAccount:  222233334444,
		PinAccount: 123456,
		Name:       "Akun2",
		Balance:    decimal.RequireFromString("1000000"),
		History: []HistoryTransaction{
			HistoryTransaction{
				Date:        time.Date(2022, 04, 26, 8, 00, 00, 00, time.Local),
				Transaction: "Setor",
				Amount:      decimal.RequireFromString("1000000"),
				Balance:     decimal.RequireFromString("1000000"),
			},
		},
	},
	AccountBank{
		IdAccount:  123412341234,
		PinAccount: 123456,
		Name:       "Akun3",
		Balance:    decimal.RequireFromString("1500000"),
		History: []HistoryTransaction{
			HistoryTransaction{
				Date:        time.Date(2022, 04, 26, 8, 00, 00, 00, time.Local),
				Transaction: "Setor",
				Amount:      decimal.RequireFromString("1500000"),
				Balance:     decimal.RequireFromString("1500000"),
			},
		},
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// menampilkan menu akun setelah login berhasil
func OptionAccount(account *AccountBank) {
	var isLogout bool = false
	for {
		fmt.Println("\nSelamat Datang,", account.Name)
		fmt.Println("1. Tarik Uang")
		fmt.Println("2. Setor Uang")
		fmt.Println("3. Transfer")
		fmt.Println("4. Lihat Riwayat Transaksi")
		fmt.Println("0. Keluar")
		fmt.Printf("\nPilih Menu [1]/[2]/[3]/[4]/[0] : ")
		option, _ := InputScan()
		switch option {
		case "0":
			fmt.Println(account.Name, "telah keluar")
			fmt.Println("Terima Kasih")
			isLogout = true
		case "1":
			tarikUang()
		case "2":
			setorUang()
		case "3":
			transferUang()
		case "4":
			riwayatTransaksi()
		default:
			fmt.Println("\n>>>>>> Pilihan Salah ! <<<<<<")
		}

		if isLogout {
			break
		}
	}
}

// fungsi untuk proses tarik uang
func tarikUang() {
	fmt.Println("Tarik Uang")
}

// fungsi untuk proses setor uang
func setorUang() {
	fmt.Println("Setor Uang")
}

// proses transfer uang
func transferUang() {
	fmt.Println("Transfer")
}

// lihat daftar riwayat transaksi
func riwayatTransaksi() {
	fmt.Println("Lihar Riwayat Transaksi")
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// menampilkan menu login
func OptionLogin() {
	fmt.Println("\n1. Login")
	fmt.Println("0. Keluar")
	fmt.Printf("\nPilih Menu [1]/[0] : ")
	option, _ := InputScan()
	switch option {
	case "0":
		fmt.Println("\n>>>>>> Terima Kasih ! <<<<<<")
		fmt.Println()
		os.Exit(0)
	case "1":
		isAccount := LoginProccess()
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

// melakukan proses login
func LoginProccess() *AccountBank {
	fmt.Println("\nSilahakan Login !")
	fmt.Println("\nMasukan Nomor Rekening :")

	var account *AccountBank
	strId, err := InputScan()
	CheckError(err)

	if cekId, intId := checkIdAccount(strId); cekId {
		// fmt.Println("\nNomor Rekening Benar !")
		fmt.Println("\nMasukan PIN :")
		strPass, err := InputScan()
		CheckError(err)
		if cekPin, cekAccount := checkPassword(strPass, intId); cekPin {
			account = cekAccount
		} else {
			fmt.Println("\nPin Anda Salah !")
		}
	} else {
		fmt.Println("\nNomor Rekening Salah !")
	}

	return account
}

// proses cek password
func checkPassword(strPass string, id int) (bool, *AccountBank) {
	var account AccountBank
	var check bool = false
	pinAccount, err := strconv.Atoi(strPass)
	if err != nil {
		return check, nil
	}
	for i, l := range listAccount {
		if l.IdAccount == id && l.PinAccount == pinAccount {
			check = true
			account = listAccount[i]
		}
	}
	return check, &account
}

// proses cek nomor rekening
func checkIdAccount(str string) (bool, int) {
	var check bool = false
	idAccount, err := strconv.Atoi(str)
	if err != nil {
		return check, idAccount
	}
	for _, l := range listAccount {
		if l.IdAccount == idAccount {
			check = true
			break
		}
	}
	return check, idAccount
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// fungsi untuk menerima input keyboard
func InputScan() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return str, err
	}
	str = strings.ReplaceAll(str, " ", "")
	str = strings.Replace(str, "\n", "", 1)
	return str, nil
}

// fungsi tambahan untuk cek error misal perlu
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error :", err.Error())
		return
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
