package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"strings"
	"time"
)

const num = 50000000
const testString = "Hello, world!"

type Logger struct{}

func (l *Logger) Info(message string) string {
	currentTime := time.Now().Format("02.01.2006")
	var sb strings.Builder
	sb.WriteString(currentTime)
	sb.WriteString(" ")
	sb.WriteString(message)
	return sb.String()
}

func main() {
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()
	log := Logger{}

	start := time.Now()

	var loggedString string
	for range num {
		loggedString = log.Info(testString)
	}
	fmt.Println(loggedString)

	elapsed := time.Since(start)

	fmt.Printf("Время выполнения (быстрое): %s\n", elapsed)

}
