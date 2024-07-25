package main

import (
	"fmt"
	"net/http"
	"time"
)

// RateLimiter для управления одновременными запросами
type RateLimiter struct {
	semaphore chan struct{}
}

// NewRateLimiter создает новый RateLimiter с заданным количеством одновременных запросов
func NewRateLimiter(maxConcurrent int) *RateLimiter {
	return &RateLimiter{
		semaphore: make(chan struct{}, maxConcurrent),
	}
}

// Allow позволяет запросу пройти, если это возможно
func (rl *RateLimiter) Allow() bool {
	select {
	case rl.semaphore <- struct{}{}:
		return true
	default:
		return false
	}
}

// Done освобождает место для следующего запроса
func (rl *RateLimiter) Done() {
	<-rl.semaphore
}

// Обработчик для HTTP-запросов
func requestHandler(w http.ResponseWriter, r *http.Request, limiter *RateLimiter) {
	if limiter.Allow() {
		defer limiter.Done() // Освобождаем место по завершении

		fmt.Fprintf(w, "Обработка запроса: %s\n", r.URL.Path)
		time.Sleep(1 * time.Second) // Симуляция обработки
	} else {
		http.Error(w, "Слишком много одновременно обрабатываемых запросов", http.StatusTooManyRequests)
	}
}

func main() {
	// Создаем rate limiter, позволяющий максимум 3 одновременных запроса
	limiter := NewRateLimiter(3)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestHandler(w, r, limiter)
	})

	fmt.Println("Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
