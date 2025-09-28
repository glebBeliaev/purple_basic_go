package main

import (
	"fmt"
	"purple_basic_go/password/account"
	"purple_basic_go/password/files"
	"purple_basic_go/password/output"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("_____Password Manager_____")
	fmt.Println(" ")
	vault := account.NewVault(files.NewJsonDb("password/data.json"))
Menu:
	for {
		variant := promtData([]string{
			"1 - Создать аккаунт",
			"2 - Найти аккаунт",
			"3 - Удалить аккаунт",
			"4 - Выход",
			"Выберите вариант"})
		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант:")
	fmt.Println("1 - Создать аккаунт")
	fmt.Println("2 - Найти аккаунт")
	fmt.Println("3 - Удалить аккаунт")
	fmt.Println("4 - Выход")
	fmt.Print("")
	fmt.Scan(&variant)
	return variant
}

func findAccount(vault *account.VaultWithDb) {
	url := promtData([]string{"Введите url для поиска"})
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		output.PrintError("Ничего не нашлось")
	}
	for _, account := range accounts {
		account.OutputData()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promtData([]string{"Введите url для поиска"})
	isDeleted := vault.DeleteAccountsByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Ничего не нашлось")
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promtData([]string{"Введите логин"})
	password := promtData([]string{"Введите пароль"})
	url := promtData([]string{"Введите url"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Не верный формат")
		return
	}
	vault.AddAccount(*myAccount)
}

func promtData[T any](promt []T) string {
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
