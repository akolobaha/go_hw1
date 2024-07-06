package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"testing"
	"time"
)

func TestSuccessResults(t *testing.T) {
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

	SendMsg("correctToken1", "test message", &wg)
	SendMsg("correctToken1", "file2", &wg)
	SendMsg("correctToken3", "test message 42", &wg)
	SendMsg("correctToken3", "test message 255", &wg)
	SendMsg("correctToken1", "file2", &wg)
	SendMsg("correctToken1", "test message", &wg)
	SendMsg("correctToken2", "-42 message", &wg)
	SendMsg("correctToken2", "test message", &wg)
	SendMsg("correctToken1", "test message", &wg)

	wg.Wait()
}

func TestWrongTokens(t *testing.T) {
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

	SendMsg("safasf", "test message", &wg)
	SendMsg("asfasf", "file2", &wg)
	SendMsg("vdzvsdv", "test message 42", &wg)
	SendMsg("asfasf", "test message 255", &wg)
	SendMsg("asdfasf", "file2", &wg)
	SendMsg("asfasdf", "test message", &wg)
	SendMsg("asfdasf", "-42 message", &wg)
	SendMsg("asdfasdf", "test message", &wg)
	SendMsg("asdfasdfasdf", "test message", &wg)

	wg.Wait()
}

func generateMD5Hash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

func TestHighLoad(t *testing.T) {
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

	for i := range 100000000 {
		tokenIndex := i % len(ValidTokens)
		SendMsg(ValidTokens[tokenIndex], generateMD5Hash(time.Now().String()), &wg)

	}

	SendMsg("correctToken1", "test message", &wg)
	SendMsg("correctToken1", "file2", &wg)
	SendMsg("correctToken3", "test message 42", &wg)
	SendMsg("correctToken3", "test message 255", &wg)
	SendMsg("correctToken1", "file2", &wg)
	SendMsg("correctToken1", "test message", &wg)
	SendMsg("correctToken2", "-42 message", &wg)
	SendMsg("correctToken2", "test message", &wg)
	SendMsg("correctToken1", "test message", &wg)

	wg.Wait()
}
