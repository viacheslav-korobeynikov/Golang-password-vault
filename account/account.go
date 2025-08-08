package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

// Объявление структуры
type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimestamp struct {
	Account
	createdAt time.Time
	updatedAt time.Time
}

// Метод структуры
func (acc AccountWithTimestamp) OutputPassword() {
	fmt.Println(acc.login, acc.password, acc.url, acc.createdAt, acc.updatedAt)
}

// Метод структуры
func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimestamp, error) {

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &AccountWithTimestamp{
		Account: Account{
			url:      urlString,
			login:    login,
			password: password,
		},
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
