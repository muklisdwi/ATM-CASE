package service

import (
	"atmcase/model"
	"atmcase/utils"
	"fmt"
	"strconv"
)

// melakukan proses login
func LoginProccess(listAccount []*model.AccountBank) *model.AccountBank {
	fmt.Println("\nMasukan Nomor Rekening :")

	var account *model.AccountBank
	strId, _ := utils.InputScan()
	if cekId, intId := checkIdAccount(strId, listAccount); cekId {
		// fmt.Println("\nNomor Rekening Benar !")
		fmt.Println("\nMasukan PIN :")
		strPass, _ := utils.InputScan()
		if cekPin, cekAccount := checkPassword(strPass, intId, listAccount); cekPin {
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
func checkPassword(strPass string, id int, listAccount []*model.AccountBank) (bool, *model.AccountBank) {
	var account *model.AccountBank
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
	return check, account
}

// proses cek nomor rekening
func checkIdAccount(str string, listAccount []*model.AccountBank) (bool, int) {
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
