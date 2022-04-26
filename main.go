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
	Date        string
	Transaction string
	Amount      decimal.Decimal
	Balance     decimal.Decimal
}

var listAccount = []AccountBank{
	{
		IdAccount:  111122223333,
		PinAccount: 123456,
		Name:       "Akun1",
		Balance:    decimal.RequireFromString("500000"),
		History: []HistoryTransaction{
			{
				Date:        TimeDateNow(),
				Transaction: "Setor",
				Amount:      decimal.RequireFromString("500000"),
				Balance:     decimal.RequireFromString("500000"),
			},
		},
	},
	{
		IdAccount:  222233334444,
		PinAccount: 123456,
		Name:       "Akun2",
		Balance:    decimal.RequireFromString("1000000"),
		History: []HistoryTransaction{
			{
				Date:        TimeDateNow(),
				Transaction: "Setor",
				Amount:      decimal.RequireFromString("1000000"),
				Balance:     decimal.RequireFromString("1000000"),
			},
		},
	},
	{
		IdAccount:  123412341234,
		PinAccount: 123456,
		Name:       "Akun3",
		Balance:    decimal.RequireFromString("1500000"),
		History: []HistoryTransaction{
			{
				Date:        TimeDateNow(),
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
			fmt.Printf("\n")
			fmt.Println(account.Name, "telah keluar")
			fmt.Println("Terima Kasih")
			isLogout = true
		case "1":
			tarikUang(account)
		case "2":
			setorUang(account)
		case "3":
			transferUang()
		case "4":
			riwayatTransaksi(account)
		default:
			fmt.Println("\n>>>>>> Pilihan Salah ! <<<<<<")
		}

		if isLogout {
			break
		}
	}
}

// fungsi untuk proses tarik uang
func tarikUang(account *AccountBank) {
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
		option, _ := InputScan()
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
			if validasiSaldo(account, tarik) {
				fmt.Println("\nAnda akan melakukan tarik uang", tarik)
				fmt.Printf("\n[1] Ya / [0] Tidak : ")
				opsi, _ := InputScan()
				if opsi == "1" {
					fmt.Println("\n>>>>> Penarikan diproses <<<<")
					result := account.Balance.Sub(decimal.NewFromInt(int64(tarik)))
					account.Balance = result
					account.History = append(account.History, HistoryTransaction{
						Date:        TimeDateNow(),
						Transaction: "Tarik",
						Amount:      decimal.NewFromInt(int64(tarik)),
						Balance:     account.Balance,
					})
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

// fungsi untuk proses setor uang
func setorUang(account *AccountBank) {
	fmt.Println("\nSetor Uang :")
	fmt.Println("Setor hanya menerima kelipatan 50000")
	fmt.Printf("\nMasukan jumlah nominal : ")
	nominal, _ := InputScan()
	setor, err := decimal.NewFromString(nominal)
	if err != nil {
		fmt.Println("\n>>>>>> Masukan Salah ! <<<<<<")
		return
	}
	if setor.Mod(decimal.NewFromInt(int64(50000))).Equal(decimal.NewFromInt32(0)) {
		fmt.Println("\n>>>>>> Setoran diproses <<<<<")
		result := setor.Add(account.Balance)
		account.Balance = result
		account.History = append(account.History, HistoryTransaction{
			Date:        TimeDateNow(),
			Transaction: "Setor",
			Amount:      setor,
			Balance:     result,
		})
	} else {
		fmt.Println("\n>>>>>> Nominal Salah ! <<<<<<")
		return
	}
}

// proses transfer uang
func transferUang() {
	fmt.Println("Transfer")
}

// lihat daftar riwayat transaksi
func riwayatTransaksi(account *AccountBank) {
	fmt.Println("\nDaftar Riwayat Transaksi :")
	for _, l := range account.History {
		fmt.Printf("%v | %s | %v | %v \n", l.Date, l.Transaction, l.Amount, l.Balance)
	}
	fmt.Printf("\n(Tekan Enter)")
	InputScan()
}

// fungsi untuk validasi saldo
func validasiSaldo(saldo *AccountBank, tarik int) bool {
	return saldo.Balance.GreaterThanOrEqual(decimal.NewFromInt(int64(tarik)))
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// menampilkan menu login
func OptionLogin() {
	fmt.Println("\nSilahakan Login !")
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

// fungsi convert time.Time
func TimeDateNow() string {
	now := time.Now()
	formatLayout := "2006-01-02 15:04:05"
	dateNow := now.Format(formatLayout)
	return dateNow
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
