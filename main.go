package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/viacheslav-korobeynikov/Golang-password-vault/account"
	"github.com/viacheslav-korobeynikov/Golang-password-vault/files"
	"github.com/viacheslav-korobeynikov/Golang-password-vault/output"
)

var menu = map[string]func(*account.VaultWithDB){
	"1": createAccount,
	"2": findAccountByLogin,
	"3": findAccountByUrl,
	"4": deleteAccountByUrl,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по логину",
	"3. Найти аккаунт по URL",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите нужный пункт меню",
}

func main() {
	fmt.Println("_Менеджер паролей_")
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось загрузить env файл")
	}
	vault := account.NewVault(files.NewJsonDB("data.json"))
Menu:
	for {
		userChoice := inputData(menuVariants...)
		menuFunc := menu[userChoice]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}

}

func createAccount(vault *account.VaultWithDB) {
	login := inputData("Введите логин")
	password := inputData("Введите пароль")
	url := inputData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный логин или URL")
		return
	}

	vault.AddAccount(*myAccount)
}

func inputData(a ...string) string {
	for i, line := range a {
		if i == len(a)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

func findAccountByUrl(vault *account.VaultWithDB) {
	url := inputData("Введите URL для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDB) {
	login := inputData("Введите логин для поиска")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}
	for _, account := range *accounts {
		account.Output()
	}
}

func deleteAccountByUrl(vault *account.VaultWithDB) {
	url := inputData("Введите URL для поиска")
	isDeleted := vault.DeleteAccountsByUrls(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Запись для удаления не найдена")
	}
}
