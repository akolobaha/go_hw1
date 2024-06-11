package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	fmt.Println("Press Ctrl+C to stop")

	sig := <-signals
	fmt.Println("Received signal:", sig)
}
