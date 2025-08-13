package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/viacheslav-korobeynikov/Golang-password-vault/account"
	"github.com/viacheslav-korobeynikov/Golang-password-vault/files"
	"github.com/viacheslav-korobeynikov/Golang-password-vault/output"
)

/*
Меню:
1. Создать аккаунт
2. Найти аккаунт
3. Удалить аккаунт
4. Выход
*/

func main() {
	fmt.Println("_Менеджер паролей_")
	vault := account.NewVault(files.NewJsonDB("data.json"))
Menu:
	for {
		userChoice := inputData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите нужный пункт меню",
		})
		switch userChoice {
		case "1":
			createAccount(vault)
		case "2":
			findAccountByUrl(vault)
		case "3":
			deleteAccountByUrl(vault)
		default:
			break Menu
		}
	}

}

func createAccount(vault *account.VaultWithDB) {
	login := inputData([]string{"Введите логин"})
	password := inputData([]string{"Введите пароль"})
	url := inputData([]string{"Введите URL"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный логин или URL")
		return
	}

	vault.AddAccount(*myAccount)
}

func inputData[T any](a []T) string {
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
	url := inputData([]string{"Введите URL для поиска"})
	accounts := vault.FindAccountsByUrls(url)
	if len(accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccountByUrl(vault *account.VaultWithDB) {
	url := inputData([]string{"Введите URL для поиска"})
	isDeleted := vault.DeleteAccountsByUrls(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Запись для удаления не найдена")
	}
}
