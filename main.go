package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

// Объявление структуры
type account struct {
	login    string
	password string
	url      string
}

// Метод структуры
func (acc account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

// Метод структуры
func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

// Функция конструктора структуры

// 1. Если логина нет, то ошибка
// 2. Если нет пароля, то генерим
func newAccount(login, password, urlString string) (*account, error) {

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	//if len(password) == 0 {
	//	(*account).generatePassword(&account.password, 12)
	//}

	newAcc := &account{
		url:      urlString,
		login:    login,
		password: password,
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func main() {

	login := inputData("Введите логин")
	password := inputData("Введите пароль")
	url := inputData("Введите URL")

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}

	myAccount.outputPassword()

}

func inputData(a string) string {
	fmt.Print(a + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
