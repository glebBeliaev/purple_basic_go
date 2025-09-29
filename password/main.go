package main

import (
	"fmt"
	"purple_basic_go/password/account"
	"purple_basic_go/password/files"
	"purple_basic_go/password/output"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	fmt.Println("_____Password Manager_____")
	fmt.Println(" ")
	vault := account.NewVault(files.NewJsonDb("password/data.json"))
Menu:
	for {
		variant := promtData(
			"1 - Создать аккаунт",
			"2 - Найти аккаунт по url",
			"3 - Найти аккаунт по логину",
			"4 - Удалить аккаунт",
			"5 - Выход",
			"Выберите вариант")
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promtData("Введите url для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promtData("Введите url для поиска")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		output.PrintError("Ничего не нашлось")
	}
	for _, account := range *accounts {
		account.OutputData()
	}

}

func deleteAccount(vault *account.VaultWithDb) {
	url := promtData("Введите url для поиска")
	isDeleted := vault.DeleteAccountsByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Ничего не нашлось")
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Не верный формат")
		return
	}
	vault.AddAccount(*myAccount)
}

func promtData(promt ...any) string {
	for i, line := range promt {
		if i == len(promt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var data string
	fmt.Scanln(&data)
	return data
}
