package main

import (
	"fmt"
	"purple_basic_go/password/account"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("_____Password Manager_____")
	fmt.Println(" ")
	vault := account.NewVault()
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
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

func findAccount(vault *account.Vault) {
	url := promtData("Введите url для поиска: ")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Ничего не нашлось")
	}
	for _, account := range accounts {
		account.OutputData()
	}
}

func deleteAccount(vault *account.Vault) {
	url := promtData("Введите url для поиска: ")
	isDeleted := vault.DeleteAccountsByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Ничего не нашлось")
	}
}

func createAccount(vault *account.Vault) {
	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите url: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Не верный формат")
		return
	}
	vault.AddAccount(*myAccount)
}

func promtData(promt string) string {
	var data string
	fmt.Print(promt)
	fmt.Scanln(&data)
	return data
}
