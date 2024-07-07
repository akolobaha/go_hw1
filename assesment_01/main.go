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

var WriterTickRate time.Duration = 3000 // Срабатывание в ms

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

	wg1 := sync.WaitGroup{}
	wg2 := sync.WaitGroup{}

	AddUser(User{"correctToken1", "file1.txt"})
	AddUser(User{"correctToken2", "file2.txt"})
	AddUser(User{"correctToken3", "file3.txt"})

	go WriteMsg2Cache(&wg1)

	WriterScale(ctx, &wg2)

	SendMsg("123456789", "test message", &wg1)
	SendMsg("correctToken1", "test message", &wg1)
	SendMsg("correctToken1", "file2", &wg1)
	SendMsg("correctToken2", "test message 42", &wg1)
	SendMsg("correctToken2", "test message 255", &wg1)
	SendMsg("42", "file2", &wg1)
	SendMsg("correctToken3", "!!", &wg1)
	SendMsg("correctToken3", "test message", &wg1)
	SendMsg("-42", "-42 message", &wg1)
	//SendMsg("42", "test message", &wg1)
	SendMsg("correctToken4", "test message", &wg1)
	SendMsg("correctToken3", "test message", &wg1)

	for i := range 10000 {
		tokenIndex := i % len(ValidTokens)
		SendMsg(ValidTokens[tokenIndex], generateMD5Hash(time.Now().String()), &wg1)

	}

	wg1.Wait()
	wg2.Wait()

}

func AddUser(user User) {
	if TokenIsValid(user.Token) {
		Users[user.Token] = user
	}
}

func SendMsg(token string, message string, wg1 *sync.WaitGroup) {
	go func() {
		if TokenIsValid(token) {
			user, ok := Users[token]
			if ok {
				MessagesMutex.Lock()
				message := Message{Token: user.Token, FieldID: user.File, Data: message}
				wg1.Add(1)
				Messages <- message
				MessagesMutex.Unlock()
			}
		}

		return
	}()
}

func WriteMsg2Cache(wg1 *sync.WaitGroup) {
	for msg := range Messages {
		CacheMutex.Lock()
		Cache[msg.FieldID] = append(Cache[msg.FieldID], msg.Data)
		CacheMutex.Unlock()
		wg1.Done()
	}

	return
}

func WriteItemToFile(wg *sync.WaitGroup) {

	// Балансировщик: поймать количество, если оно возрастет - замасштабировать
	CacheMutex.Lock()
	for fieldID, data := range Cache {
		file, err := os.OpenFile(fieldID, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			// fmt.Println("Error creating file:", err)
			return
		}

		for _, str := range data {
			_, err = file.WriteString(str + "\n")
			if err != nil {
				// fmt.Println("Error writing to file:", err)
				return
			}
		}

		file.Close()

		delete(Cache, fieldID)
	}
	wg.Done()
	CacheMutex.Unlock()
}

func WriterScale(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println()
	go Writer(ctx, wg)
	go Writer(ctx, wg)
	go Writer(ctx, wg)
}

func Writer(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	ticker := time.NewTicker(WriterTickRate * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			CacheMutex.RLock()
			for fieldID, data := range Cache {
				fmt.Println(fieldID, len(data))
			}
			go WriteItemToFile(wg)
			CacheMutex.RUnlock()

			return
		case <-ticker.C:
			CacheMutex.RLock()
			for fieldID, data := range Cache {
				fmt.Println(fieldID, len(data))
			}
			go WriteItemToFile(wg)
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
