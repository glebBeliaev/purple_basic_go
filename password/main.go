package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) outputData() {
	fmt.Println("Логин:", acc.login)
	fmt.Println("Пароль:", acc.password)
	fmt.Println("URL:", acc.url)
}

func (acc *account) generatePassword(n int) {
	var letterRuns = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRuns[rand.Intn(len(letterRuns))]
	}
	acc.password = string(res)
}

func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &account{
		login:    login,
		password: password,
		url:      urlString,
	}
	if newAcc.password == "" {
		newAcc.generatePassword(8)
	}
	return newAcc, nil
}

func main() {
	login := promtData("Введите логин: ")
	password := promtData("Введите пароль: ")
	url := promtData("Введите url: ")

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Не верный формат")
		return
	}
	myAccount.outputData()

}

func promtData(promt string) string {
	var data string
	fmt.Print(promt)
	fmt.Scanln(&data)
	return data
}
