package main

import (
	"fmt"
	"os"
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

var ValidTokens = []string{"255", "42", "-1", "-42"}
var Users = make(map[string]User)
var Messages = make(chan Message)
var Cache = make(map[string][]string)
var CacheMutex sync.RWMutex

var WriterTickRate time.Duration = 5000 // Срабатывание в ms

func main() {
	wg := sync.WaitGroup{}

	AddUser(User{"42", "file42.txt"})
	AddUser(User{"-42", "file-42.txt"})
	AddUser(User{"255", "file-255.txt"})

	go WriteMsg2Cache(&wg)

	go WriteFiles(&wg)

	SendMsg("123456789", "test message", &wg)
	SendMsg("42", "test message", &wg)
	SendMsg("42", "file2", &wg)
	SendMsg("255", "test message 42", &wg)
	SendMsg("255", "test message 255", &wg)
	SendMsg("42", "file2", &wg)
	SendMsg("421", "!!", &wg)
	SendMsg("42", "test message", &wg)
	SendMsg("-42", "-42 message", &wg)
	//SendMsg("42", "test message", &wg)
	SendMsg("-42", "test message", &wg)
	SendMsg("42", "test message", &wg)

	wg.Wait()

	fmt.Println(Cache)
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

func WriteFiles(wg *sync.WaitGroup) {
	wg.Add(1)
	time.Sleep(1000 * time.Millisecond)

	ticker := time.NewTicker(WriterTickRate * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
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
				fmt.Println(Cache)

				delete(Cache, fieldID)
			}
			CacheMutex.Unlock()
			wg.Done()
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
