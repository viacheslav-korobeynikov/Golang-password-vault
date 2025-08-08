package files

import (
	"fmt"
	"os"
)

func ReadFile() {
	//file, err := os.Open("file.txt") - чтение файла по байтам
	data, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

func WriteFile(content string, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(content)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись завершена успешно")
}
