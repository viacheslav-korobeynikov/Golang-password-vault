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

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func main() {
	fmt.Println(generatePassword(12))

	//fmt.Println(rand.Int32N(10))

	//str := []rune("Привет!)")
	//for _, ch := range string(str) {
	//	fmt.Println(ch, string(ch))
	//}

	login := inputData("Введите логин")
	password := inputData("Введите пароль")
	url := inputData("Введите URL")

	myAccount := account{
		login:    login,
		password: password,
		url:      url,
	}

	outputPassword(&myAccount)

}

func inputData(a string) string {
	fmt.Print(a + ": ")
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url)
}

// Функция принимает на вход количество символов в пароле и возвращает пароль
// Массив допустимых значений

func generatePassword(n int) string {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	return string(res)
}
