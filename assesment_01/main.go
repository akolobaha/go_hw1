package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"runtime/trace"
	"sync/atomic"
	"time"
)

type User struct {
	Token string
	File  string
}

type Message struct {
	Token  string
	FileID string
	Data   string
}

var ValidTokens = []string{"correctToken1", "correctToken2", "correctToken3", "correctToken4"}
var Users = make(map[string]User)
var Messages = make(chan Message, 100000)
var WorkersCount int32 = 1
var done = make(chan struct{})

var ScaleFactor int = 10          // Коэффицент масштабинования (Сообщений на воркер)
var TickRate time.Duration = 1000 // Задержка воркера ms

func main() {
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()
	debug.SetGCPercent(500)

	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		select {
		case <-quit:
			cancel()
			fmt.Println("Ctrl + C pressed, canceling context")
		}
	}()

	AddUser(User{"correctToken1", "file1.txt"})
	AddUser(User{"correctToken2", "file2.txt"})
	AddUser(User{"correctToken3", "file3.txt"})
	AddUser(User{"correctToken4", "file4.txt"})

	for i := range 100 {
		tokenIndex := i % len(ValidTokens)
		SendMsg(ValidTokens[tokenIndex], generateMD5Hash(time.Now().String()))
	}

	Worker(ctx, cancel)

	<-done
	fmt.Println("Done")
}

func Worker(ctx context.Context, cancelFunc context.CancelFunc) {
	go func() {
		ticker := time.NewTicker(TickRate * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				if len(Messages) > 0 {
					message := <-Messages
					fmt.Println("Запись в файл: ", message.Data, message.FileID)
					WriteMessageToFile(message)
				} else {
					done <- struct{}{}
					cancelFunc()
				}
			case <-ticker.C:
				currentWorkers := atomic.LoadInt32(&WorkersCount)
				requiredWorkers := len(Messages) / ScaleFactor

				if currentWorkers < int32(requiredWorkers) {
					atomic.AddInt32(&WorkersCount, 1)
					Worker(ctx, cancelFunc)
					fmt.Println("Воркер замасштабировался до ", atomic.LoadInt32(&WorkersCount))
				}

				if len(Messages) > 0 {
					message := <-Messages
					fmt.Println("Запись в файл: ", message.Data, message.FileID)
					WriteMessageToFile(message)
				} else {
					done <- struct{}{}
					cancelFunc()
				}
			}

		}
	}()
}

func AddUser(user User) {
	if TokenIsValid(user.Token) {
		Users[user.Token] = user
	}
}

func SendMsg(token string, message string) {
	go func() {
		if TokenIsValid(token) {
			user, ok := Users[token]
			if ok {
				message := Message{Token: user.Token, FileID: user.File, Data: message}
				Messages <- message
			}
		}

		return
	}()
}

func WriteMessageToFile(message Message) {
	file, err := os.OpenFile(message.FileID, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}

	// retry
	var writeSuccess bool
	for i := 0; i < 10; i++ {
		_, err = file.WriteString(message.Data + "\n")
		if err == nil {
			writeSuccess = true
			break
		}
		fmt.Println("Конфликт при записи")
		time.Sleep(10 * time.Millisecond)
	}
	if !writeSuccess {
		return
	}

	file.Close()
}

func TokenIsValid(token string) bool {
	for _, validToken := range ValidTokens {
		if token == validToken {
			return true
		}
	}
	return false
}

/*
Моковая строка
*/
func generateMD5Hash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}
