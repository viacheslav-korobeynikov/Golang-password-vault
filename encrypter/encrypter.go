package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewRncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан параметр KEY в переменные окружения")
	}
	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encrypt(cleanStr []byte) []byte {
	// Создаем блок для шифрования
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	// Формируем 128-битный блочный шифр, упакованный в режим счетчика Галуа
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	// Некоторое одноразование число, которое добавляется к блоку
	nonce := make([]byte, aesGSM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	// Шифруем (через функцию Seal) и возвращаем зашифрованный массив байт
	return aesGSM.Seal(nonce, nonce, cleanStr, nil)
}

func (enc *Encrypter) Decrypt(encryptedStr []byte) []byte {
	// Создаем блок для шифрования
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	// Формируем 128-битный блочный шифр, упакованный в режим счетчика Галуа
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	NonceSize := aesGSM.NonceSize()
	nonce := encryptedStr[:NonceSize]
	cipherText := encryptedStr[NonceSize:]
	// Дешифруем (через функцию Open)
	cleanStr, err := aesGSM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	// Возвращаем дешифрованный массив байт
	return cleanStr
}
