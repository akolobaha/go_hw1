package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	create := flag.String("create", "", "Создать файл")
	read := flag.String("read", "", "Прочитать содержимое файла")
	remove := flag.String("delete", "", "Удалить файл")

	flag.Parse()

	createByFileName(*create)
	readByFileName(*read)
	deleteByFileName(*remove)
}

func createByFileName(fileName string) {
	if len(fileName) > 0 {
		_, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Ошибка создания файла:", err)
			return
		}
	}
}

func readByFileName(fileName string) {
	if len(fileName) > 0 {
		data, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Ошибка чтения файла:", err)
			return
		}

		fmt.Println(string(data))
	}
}

func deleteByFileName(fileName string) {
	if len(fileName) > 0 {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println("Ошибка удаления файла:", err)
		}
	}
}
