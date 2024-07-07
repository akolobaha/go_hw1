package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
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

	done := make(chan struct{})

	AddUser(User{"correctToken1", "file1.txt"})
	AddUser(User{"correctToken2", "file2.txt"})
	AddUser(User{"correctToken3", "file3.txt"})

	go WriteMsg2Cache()

	go Writer(ctx, done)

	SendMsg("correctToken1", "test message")
	SendMsg("correctToken1", "file2")
	SendMsg("correctToken3", "test message 42")
	SendMsg("correctToken3", "test message 255")
	SendMsg("correctToken1", "file2")
	SendMsg("correctToken1", "test message")
	SendMsg("correctToken2", "-42 message")
	SendMsg("correctToken2", "test message")
	SendMsg("correctToken1", "test message")

	<-done
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

	done := make(chan struct{})

	AddUser(User{"correctToken1", "file1.txt"})
	AddUser(User{"correctToken2", "file2.txt"})
	AddUser(User{"correctToken3", "file3.txt"})

	go WriteMsg2Cache()

	go Writer(ctx, done)

	SendMsg("safasf", "test message")
	SendMsg("asfasf", "file2")
	SendMsg("vdzvsdv", "test message 42")
	SendMsg("asfasf", "test message 255")
	SendMsg("asdfasf", "file2")
	SendMsg("asfasdf", "test message")
	SendMsg("asfdasf", "-42 message")
	SendMsg("asdfasdf", "test message")
	SendMsg("asdfasdfasdf", "test message")

	<-done
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

	done := make(chan struct{})

	AddUser(User{"correctToken1", "file1.txt"})
	AddUser(User{"correctToken2", "file2.txt"})
	AddUser(User{"correctToken3", "file3.txt"})

	go WriteMsg2Cache()

	go Writer(ctx, done)

	for i := range 1000000 {
		tokenIndex := i % len(ValidTokens)
		SendMsg(ValidTokens[tokenIndex], generateMD5Hash(time.Now().String()))

	}

	SendMsg("correctToken1", "test message")
	SendMsg("correctToken1", "file2")
	SendMsg("correctToken3", "test message 42")
	SendMsg("correctToken3", "test message 255")
	SendMsg("correctToken1", "file2")
	SendMsg("correctToken1", "test message")
	SendMsg("correctToken2", "-42 message")
	SendMsg("correctToken2", "test message")
	SendMsg("correctToken1", "test message")

	<-done
}
