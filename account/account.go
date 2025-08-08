package account

import (
	"encoding/json"
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
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Метод структуры
func (acc Account) OutputPassword() {
	color.Cyan(acc.Login)
	color.Magenta(acc.Password)
	color.Blue(acc.Url)
}

// Метод структуры
func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

func (acc *Account) ToByte() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewAccount(login, password, urlString string) (*Account, error) {

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{
		Url:       urlString,
		Login:     login,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	//field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login")
	//fmt.Println(string(field.Tag))
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
