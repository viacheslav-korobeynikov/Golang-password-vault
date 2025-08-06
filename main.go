package main

import "fmt"

type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := inputData("Введите логин")
	password := inputData("Введите пароль")
	url := inputData("Введите URL")

	myAccount := account{
		login:    login,
		password: password,
		url:      url,
	}

	outputPassword(myAccount)

}

func inputData(a string) string {
	fmt.Print(a + ": ")
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc account) {
	fmt.Println(acc.login, acc.password, acc.url)
}
