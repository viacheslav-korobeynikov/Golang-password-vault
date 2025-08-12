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
		userChoice := showMenu()
		switch userChoice {
		case 1:
			createAccount(vault)
		case 2:
			findAccountByUrl(vault)
		case 3:
			deleteAccountByUrl(vault)
		default:
			break Menu
		}
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

func inputData(a string) string {
	fmt.Print(a + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

func showMenu() int {
	var userChoice int
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scanln(&userChoice)
	return userChoice
}

func findAccountByUrl(vault *account.VaultWithDB) {
	url := inputData("Введите URL для поиска")
	accounts := vault.FindAccountsByUrls(url)
	if len(accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}
	for _, account := range accounts {
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
