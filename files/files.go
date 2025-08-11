package files

import (
	"fmt"
	"os"
)

type JsonDB struct {
	filename string
}

func NewJsonDB(name string) *JsonDB {
	return &JsonDB{
		filename: name,
	}
}

func (db *JsonDB) Read() ([]byte, error) {
	//file, err := os.Open("file.txt") - чтение файла по байтам
	data, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (db *JsonDB) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись завершена успешно")
}
