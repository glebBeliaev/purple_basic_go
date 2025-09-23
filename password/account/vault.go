package account

import (
	"encoding/json"
	"strings"

	"purple_basic_go/password/files"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) FindAccountsByUrl(url string) []Account {
	var accounts []Account
	for _, acc := range vault.Accounts {
		isMatched := strings.Contains(acc.Url, url)
		if isMatched {
			accounts = append(accounts, acc)
		}
	}
	return accounts
}

func (vault *Vault) DeleteAccountsByUrl(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, acc := range vault.Accounts {
		isMatched := strings.Contains(acc.Url, url)
		if !isMatched {
			accounts = append(accounts, acc)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.Save()
	return isDeleted
}

func (vault *Vault) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.Save()
}

func NewVault() *Vault {
	db := files.NewJsonDb("password/data.json")
	file, err := db.Read()
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл JSON")
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	return &vault
}

func (vault *Vault) Save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	db := files.NewJsonDb("password/data.json")
	db.Write(data)
}
