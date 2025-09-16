package account

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"

	"github.com/fatih/color"
)

type Account struct {
	login    string
	password string
	url      string
}

func (acc *Account) OutputData() {
	color.Green("Логин: %s", acc.login)
	fmt.Println("Пароль:", acc.password)
	fmt.Println("URL:", acc.url)
}

func (acc *Account) generatePassword(n int) {
	var letterRuns = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRuns[rand.Intn(len(letterRuns))]
	}
	acc.password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &Account{
		login:    login,
		password: password,
		url:      urlString,
	}
	if newAcc.password == "" {
		newAcc.generatePassword(8)
	}
	return newAcc, nil
}
