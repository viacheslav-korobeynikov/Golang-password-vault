package account

import (
	"encoding/json"
	"github.com/fatih/color"
	"strings"
	"time"

	"github.com/viacheslav-korobeynikov/Golang-password-vault/encrypter"
	"github.com/viacheslav-korobeynikov/Golang-password-vault/output"
)

type DB interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDB struct {
	Vault
	db  DB
	enc encrypter.Encrypter
}

func NewVault(db DB, enc encrypter.Encrypter) *VaultWithDB {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	decryptedData := enc.Decrypt(file)
	var vault Vault
	err = json.Unmarshal(decryptedData, &vault)
	color.Yellow("Найдено %d аккаунтов", len(vault.Accounts))
	if err != nil {
		output.PrintError("Не удалось разобрать файл data.vault")
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	return &VaultWithDB{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultWithDB) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := checker(account, str)
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
	encData := vault.enc.Encrypt(data)
	if err != nil {
		output.PrintError("Не удалось преобразовать")
	}
	vault.db.Write(encData)
}
