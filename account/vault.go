package account

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/viacheslav-korobeynikov/Golang-password-vault/files"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл")
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	return &vault
}

func (vault *Vault) FindAccountsByUrls(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *Vault) DeleteAccountsByUrls(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.saveVault()
	return isDeleted
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.saveVault()
}

func (vault *Vault) ToByte() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) saveVault() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToByte()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	files.WriteFile(data, "data.json")
}
