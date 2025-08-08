package files

import (
	"fmt"
	"os"
)

func ReadFile() {

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
