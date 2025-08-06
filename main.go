package main

import "fmt"

type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := inputData("Введите логин: ")
	password := inputData("Введите пароль: ")
	url := inputData("Введите URL: ")

	account1 := account{
		login,
		password,
		url,
	}
	account2 := account{
		login:    login,
		password: password,
		url:      url,
	}

	outputPassword(login, password, url)

}

func inputData(a string) string {
	fmt.Println(a)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(login string, password string, url string) {
	fmt.Println(login, password, url)
}
