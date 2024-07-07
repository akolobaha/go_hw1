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
	"sync"
	"time"
)

type User struct {
	Token string
	File  string
}

type Message struct {
	Token   string
	FieldID string
	Data    string
}

var ValidTokens = []string{"correctToken1", "correctToken2", "correctToken3", "correctToken4"}
var Users = make(map[string]User)
var Messages = make(chan Message, 100)
var Cache = make(map[string][]string)

var CacheMutex sync.RWMutex
var MessagesMutex sync.RWMutex

var done = make(chan struct{})

var WriterTickRate time.Duration = 5000 // Срабатывание в ms

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

	go WriteMsg2Cache()

	go WriterScale(ctx, done)

	SendMsg("123456789", "test message")
	SendMsg("correctToken1", "test message")
	SendMsg("correctToken1", "file2")
	SendMsg("correctToken2", "test message 42")
	SendMsg("correctToken2", "test message 255")
	SendMsg("42", "file2")
	SendMsg("correctToken3", "!!")
	SendMsg("correctToken3", "test message")
	SendMsg("-42", "-42 message")
	//SendMsg("42", "test message")
	SendMsg("correctToken4", "test message")
	SendMsg("correctToken3", "test message")

	for i := range 1000000 {
		tokenIndex := i % len(ValidTokens)
		SendMsg(ValidTokens[tokenIndex], generateMD5Hash(time.Now().String()))

	}

	<-done
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
				MessagesMutex.Lock()
				message := Message{Token: user.Token, FieldID: user.File, Data: message}
				Messages <- message
				MessagesMutex.Unlock()
			}
		}

		return
	}()
}

func WriteMsg2Cache() {
	for msg := range Messages {
		CacheMutex.Lock()
		Cache[msg.FieldID] = append(Cache[msg.FieldID], msg.Data)
		CacheMutex.Unlock()
	}

	return
}

func WriteItemToFile(done chan<- struct{}) {
	CacheMutex.RLock()
	for fieldID, data := range Cache {
		file, err := os.OpenFile(fieldID, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			// fmt.Println("Error creating file:", err)
			return
		}

		var writeSuccess bool
		for _, str := range data {
			for i := 0; i < 3; i++ {
				_, err = file.WriteString(str + "\n")
				if err == nil {
					writeSuccess = true
					break
				}
				time.Sleep(time.Second)
			}
			if !writeSuccess {
				return
			}
		}

		file.Close()

		delete(Cache, fieldID)
	}
	done <- struct{}{}
	CacheMutex.RLock()
}

func WriterScale(ctx context.Context, done chan<- struct{}) {
	go Writer(ctx, done)
}

func Writer(ctx context.Context, done chan<- struct{}) {
	ticker := time.NewTicker(WriterTickRate * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			CacheMutex.RLock()
			go WriteItemToFile(done)
			CacheMutex.RUnlock()
			return
		case <-ticker.C:
			CacheMutex.RLock()
			go WriteItemToFile(done)
			CacheMutex.RUnlock()
		}

	}
}

func TokenIsValid(token string) bool {
	for _, validToken := range ValidTokens {
		if token == validToken {
			return true
		}
	}
	return false
}

func generateMD5Hash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}
