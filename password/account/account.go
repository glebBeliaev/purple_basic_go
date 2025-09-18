package account

import (
	"errors"
	"math/rand"
	"net/url"

	"github.com/fatih/color"
)

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

func (acc *Account) OutputData() {
	color.Green("Логин: %s", acc.Login)
	color.Green("Пароль: %s", acc.Password)
	color.Green("URL: %s", acc.Url)
}

func (acc *Account) generatePassword(n int) {
	var letterRuns = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRuns[rand.Intn(len(letterRuns))]
	}
	acc.Password = string(res)
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
		Login:    login,
		Password: password,
		Url:      urlString,
	}
	if newAcc.Password == "" {
		newAcc.generatePassword(8)
	}
	return newAcc, nil
}
