package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
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
var Messages = make(chan Message)
var Cache = make(map[string][]string)
var CacheMutex sync.RWMutex

var WriterTickRate time.Duration = 5000 // Срабатывание в ms

func main() {

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

	wg := sync.WaitGroup{}

	AddUser(User{"correctToken1", "file1.txt"})
	AddUser(User{"correctToken2", "file2.txt"})
	AddUser(User{"correctToken3", "file3.txt"})

	go WriteMsg2Cache(&wg)

	go WriteFiles(ctx, &wg)

	SendMsg("123456789", "test message", &wg)
	SendMsg("correctToken1", "test message", &wg)
	SendMsg("correctToken1", "file2", &wg)
	SendMsg("correctToken2", "test message 42", &wg)
	SendMsg("correctToken2", "test message 255", &wg)
	SendMsg("42", "file2", &wg)
	SendMsg("correctToken3", "!!", &wg)
	SendMsg("correctToken3", "test message", &wg)
	SendMsg("-42", "-42 message", &wg)
	//SendMsg("42", "test message", &wg)
	SendMsg("correctToken4", "test message", &wg)
	SendMsg("42", "test message", &wg)

	wg.Wait()

}

func AddUser(user User) {
	if TokenIsValid(user.Token) {
		Users[user.Token] = user
	}
}

func SendMsg(token string, message string, wg *sync.WaitGroup) {
	if TokenIsValid(token) {
		user, ok := Users[token]
		if ok {
			message := Message{Token: user.Token, FieldID: user.File, Data: message}
			Messages <- message
			wg.Add(1)
		}
	} else {
		fmt.Println("SendMsg: Token is invalid")
	}
	return
}

func WriteMsg2Cache(wg *sync.WaitGroup) {
	for msg := range Messages {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(msg.Token, msg.FieldID, msg.Data)

		CacheMutex.Lock()
		Cache[msg.FieldID] = append(Cache[msg.FieldID], msg.Data)
		CacheMutex.Unlock()

		wg.Done()
	}

	return
}

func WriteItemToFile(wg *sync.WaitGroup) {
	CacheMutex.Lock()
	for fieldID, data := range Cache {
		file, err := os.OpenFile(fieldID, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		for _, str := range data {
			_, err = file.WriteString(str + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}

		file.Close()

		delete(Cache, fieldID)
	}
	CacheMutex.Unlock()
	wg.Done()
}

func WriteFiles(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	time.Sleep(1000 * time.Millisecond)

	ticker := time.NewTicker(WriterTickRate * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			WriteItemToFile(wg)
			return
		case <-ticker.C:
			WriteItemToFile(wg)
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
