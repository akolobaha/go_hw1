package main

import (
	"fmt"
	"net/http"
	"time"
)

var url = "localhost:8080"
var limit = 10

// RateLimiter для управления одновременными запросами
type RateLimiter struct {
	semaphore chan struct{}
}

func NewRateLimiter(maxConcurrent int) *RateLimiter {
	return &RateLimiter{
		semaphore: make(chan struct{}, maxConcurrent),
	}
}

func (rl *RateLimiter) Allow() bool {
	select {
	case rl.semaphore <- struct{}{}:
		return true
	default:
		return false
	}
}

func (rl *RateLimiter) Done() {
	<-rl.semaphore
}

func requestHandler(w http.ResponseWriter, r *http.Request, limiter *RateLimiter) {
	if limiter.Allow() {
		defer limiter.Done()

		_, err := fmt.Fprintf(w, "Обработка запроса: %s\n", r.URL.Path)
		if err != nil {
			return
		}
		time.Sleep(1500 * time.Millisecond) // Симуляция обработки
	} else {
		http.Error(w, "Слишком много одновременно обрабатываемых запросов", http.StatusTooManyRequests)
	}
}

func main() {
	limiter := NewRateLimiter(limit)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestHandler(w, r, limiter)
	})

	fmt.Println("Сервер запущен на http://", url)
	if err := http.ListenAndServe(url, nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
