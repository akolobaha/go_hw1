package main

import (
	"bufio"
	"fmt"
	"os"
)

const outputFilename = "output.txt"

func main() {
	chn := processInput()
	done := addToFile(chn)
	<-done
}

func processInput() chan string {
	ch := make(chan string)

	go func() {
		for {
			var input string
			_, err := fmt.Scanln(&input)
			if err != nil {
				fmt.Println("Ошибка сканирования строки:", err)
				return
			}

			ch <- input
		}
	}()

	return ch
}

func addToFile(ch <-chan string) chan struct{} {
	done := make(chan struct{})

	// Создаем файл для записи
	file, err := os.Create(outputFilename)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
	}

	// Создаем объект писателя
	writer := bufio.NewWriter(file)

	go func() {
		defer close(done)
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)

		for text := range ch {
			// Записываем строку в файл
			_, err = writer.WriteString(text + "\n")
			if err != nil {
				fmt.Println("Ошибка при записи в файл:", err)
				return
			}

			// Сбрасываем буфер и записываем данные на диск
			err = writer.Flush()
			if err != nil {
				fmt.Println("Ошибка при сбросе буфера:", err)
				return
			}

		}
	}()

	return done
}
