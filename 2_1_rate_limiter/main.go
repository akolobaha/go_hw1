package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const maxConnections = 5

var (
	currentConnections int
	mutex              sync.Mutex
)

func handler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	if currentConnections >= maxConnections {
		mutex.Unlock()
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	currentConnections++
	time.Sleep(1 * time.Second)
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		currentConnections--
		mutex.Unlock()
	}()

	// Обработка запроса
	time.Sleep(1 * time.Second)
	fmt.Fprintf(w, "Welcome! Current connections: %d", currentConnections)
}

func main() {

	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}

	client()
}

// var url = "localhost:8080"
//
// const maxConnections = 10
//
// var sem = make(chan struct{}, maxConnections)
//
//	func handler(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "Hello World")
//	}
//
//	func main() {
//		go server()
//
//		client()
//	}
//
//	func server() {
//		http.HandleFunc("/", handler)
//
//		fmt.Println("Сервер запущен на http://localhost:8080/")
//		if err := http.ListenAndServe(url, nil); err != nil {
//			fmt.Println("Ошибка при запуске сервера:", err)
//		}
//	}
func client() {
	for range 100 {
		func() {
			resp, err := http.Get("http://" + "localhost:8080")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
			}
			fmt.Println(resp.StatusCode, string(body))
		}()

	}
}
