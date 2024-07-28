package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

const num = 50000000
const testString = "Hello, world!"

type Logger struct{}

func (l *Logger) Info(message string) *string {
	now := time.Now()
	formattedTime := now.Format("02.01.2006")

	fmtString := fmt.Sprintf("%s: %s", formattedTime, message)
	return &fmtString
}

func main() {
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()
	log := Logger{}

	start := time.Now()

	var loggedString string
	for range num {
		loggedString = *log.Info(testString)
	}
	fmt.Println(loggedString)

	elapsed := time.Since(start)

	fmt.Printf("Время выполнения (медленное): %s\n", elapsed)

}
