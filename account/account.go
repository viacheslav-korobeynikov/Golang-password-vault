package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"reflect"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

// Объявление структуры
type Account struct {
	login    string `json:"login" xml: "test"`
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
	color.Cyan(acc.login)
	color.Magenta(acc.password)
	color.Blue(acc.url)
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
	field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login")
	fmt.Println(string(field.Tag))
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
