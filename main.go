package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func main() {

	login := inputData("Введите логин")
	url := inputData("Введите URL")

	myAccount := account{
		login: login,
		url:   url,
	}
	myAccount.generatePassword(12)
	myAccount.outputPassword()

}

func inputData(a string) string {
	fmt.Print(a + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
