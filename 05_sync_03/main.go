package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

const outputFilename = "output.txt"

func main() {
	ch := make(chan string)
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)

	processInput(cancel, ch)
	addToFile(ctx, ch, wg)

	wg.Wait()
}

func processInput(cancel context.CancelFunc, ch chan<- string) {
	fmt.Println("Press Ctrl+C to stop")
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		defer signal.Stop(done)
		defer cancel()
		<-done
	}()

	go func() {
		var result string
		for {
			_, err := fmt.Scanln(&result)
			if err != nil {
				return
			}
			ch <- result
		}
	}()
}

func addToFile(ctx context.Context, ch <-chan string, wg *sync.WaitGroup) {
	// Создаем файл для записи
	file, err := os.Create(outputFilename)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
	}
	// Создаем объект писателя
	writer := bufio.NewWriter(file)

	go func() {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			case text := <-ch:
				_, err = writer.WriteString(text + "\n")
				if err != nil {
					fmt.Println("Ошибка при записи в файл:", err)
					return
				}

				err = writer.Flush()
				if err != nil {
					fmt.Println("Ошибка при сбросе буфера:", err)
					return
				}
			}
		}
	}()
}
