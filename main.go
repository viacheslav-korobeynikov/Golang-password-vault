package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/viacheslav-korobeynikov/Golang-password-vault/account"
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
	vault := account.NewVault()
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

func createAccount(vault *account.Vault) {
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

func findAccountByUrl(vault *account.Vault) {
	url := inputData("Введите URL для поиска")
	accounts := vault.FindAccountsByUrls(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccountByUrl(vault *account.Vault) {
	url := inputData("Введите URL для поиска")
	isDeleted := vault.DeleteAccountsByUrls(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Запись для удаления не найдена")
	}
}
