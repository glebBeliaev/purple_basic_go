package main

import (
	"fmt"
	"purple_basic_go/password/account"
)

func main() {
	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите url: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Не верный формат")
		return
	}
	myAccount.OutputData()

}

func promtData(promt string) string {
	var data string
	fmt.Print(promt)
	fmt.Scanln(&data)
	return data
}
