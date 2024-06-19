package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	ch := make(chan string)
	go printInput(ch)

	readInput(ch)
}

func readInput(ch chan string) {
	var result string
	for {
		_, err := fmt.Scanln(&result)
		if err != nil {
			return
		}

		ch <- result
	}
}

func printInput(ch chan string) chan bool {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	fmt.Println("Press Ctrl+C to stop")

	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		}
	}
}
