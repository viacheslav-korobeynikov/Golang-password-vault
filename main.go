package main

import (
	"fmt"
	"github.com/viacheslav-korobeynikov/Golang-password-vault/account"
)

func main() {

	login := inputData("Введите логин")
	password := inputData("Введите пароль")
	url := inputData("Введите URL")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}

	myAccount.OutputPassword()

}

func inputData(a string) string {
	fmt.Print(a + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
