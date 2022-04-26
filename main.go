package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

type AccountBank struct {
	IdAccount  int
	PinAccount int
	Name       string
	Balance    decimal.Decimal
}

var listAccount = []AccountBank{
	AccountBank{
		IdAccount:  111122223333,
		PinAccount: 123456,
		Name:       "Akun1",
		Balance:    decimal.Decimal{},
	},
	AccountBank{
		IdAccount:  222233334444,
		PinAccount: 123456,
		Name:       "Akun2",
		Balance:    decimal.Decimal{},
	},
	AccountBank{
		IdAccount:  123412341234,
		PinAccount: 123456,
		Name:       "Akun3",
		Balance:    decimal.Decimal{},
	},
}

func main() {
	StartApp()
}

func StartApp() {
	fmt.Println("\n>>>>>> Selamat Datang ! <<<<<<")
	for {
		OptionLogin()
		// fmt.Println(LoginProccess())
	}
}

func OptionLogin() {
	fmt.Println("\n1. Login")
	fmt.Println("0. Keluar")
	fmt.Printf("\nPilih Menu [0]/[1] : ")
	option, _ := InputScan()
	switch option {
	case "0":
		fmt.Println("\n>>>>>> Terima Kasih ! <<<<<<")
		fmt.Println()
		os.Exit(0)
	case "1":
		isLogin := LoginProccess()
		if isLogin != nil {
			fmt.Println(isLogin)
		} else {
			fmt.Println("Login Gagal !")
		}
	default:
		fmt.Println("\n>>>>>> Pilihan Salah ! <<<<<<")
		fmt.Println()
	}
}

func LoginProccess() *AccountBank {
	fmt.Println("\nSilahakan Login !")
	fmt.Println("\nMasukan Nomor Rekening :")

	var account *AccountBank
	strId, err := InputScan()
	CheckError(err)

	if cekId, intId := CheckIdAccount(strId); cekId {
		// fmt.Println("\nNomor Rekening Benar !")
		fmt.Println("\nMasukan PIN :")
		strPass, err := InputScan()
		CheckError(err)
		if cekPin, cekAccount := CheckPassword(strPass, intId); cekPin {
			account = cekAccount
		} else {
			fmt.Println("\nPin Anda Salah !")
		}
	} else {
		fmt.Println("\nNomor Rekening Salah !")
	}

	return account
}

func CheckPassword(strPass string, id int) (bool, *AccountBank) {
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

func CheckIdAccount(str string) (bool, int) {
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

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error :", err.Error())
		return
	}
}
