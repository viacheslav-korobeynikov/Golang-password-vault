package encrypter

import "os"

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

func (enc *Encrypter) Encrypt(cleanStr string) string {
	return ""
}

func (enc *Encrypter) Decrypt(encryptedStr string) string {
	return ""
}
