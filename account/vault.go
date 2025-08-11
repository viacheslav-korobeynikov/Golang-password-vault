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

type VaultWithDB struct {
	Vault
	db files.JsonDB
}

func NewVault(db *files.JsonDB) *VaultWithDB {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: *db,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл")
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: *db,
		}
	}
	return &VaultWithDB{
		Vault: vault,
		db:    *db,
	}
}

func (vault *VaultWithDB) FindAccountsByUrls(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *VaultWithDB) DeleteAccountsByUrls(url string) bool {
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

func (vault *VaultWithDB) AddAccount(acc Account) {
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

func (vault *VaultWithDB) saveVault() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToByte()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	vault.db.Write(data)
}
