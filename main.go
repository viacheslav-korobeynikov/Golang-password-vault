package main

import (
	"fmt"

	"github.com/viacheslav-korobeynikov/Golang-password-vault/account"
	"github.com/viacheslav-korobeynikov/Golang-password-vault/files"
)

/*
Меню:
1. Создать аккаунт
2. Найти аккаунт
3. Удалить аккаунт
4. Выход
*/

func main() {
	//files.ReadFile()
	//files.WriteFile("Hello, world!", "file.txt")
	fmt.Println("_Менеджер паролей_")
Menu:
	for {
		userChoice := showMenu()
		switch userChoice {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}

}

func createAccount() {
	login := inputData("Введите логин")
	password := inputData("Введите пароль")
	url := inputData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := myAccount.ToByte()
	if err != nil {
		fmt.Println("Не удалось преобразовать данные в JSON")
		return
	}
	files.WriteFile(file, "data.json")
}

func inputData(a string) string {
	fmt.Print(a + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

func showMenu() int {
	var userChoice int

	fmt.Println("Выберите нужный пункт меню: ")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scan(&userChoice)
	return userChoice
}

func findAccount() {

}

func deleteAccount() {

}
