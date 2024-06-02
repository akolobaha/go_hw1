package main

import (
	"fmt"
	"github.com/fxtlabs/primes"
)

func main() {
	primesChan, compositesChan := splitPrimes([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	done1 := goPrinter(primesChan)
	done2 := goPrinter(compositesChan)

	<-done1
	<-done2
}

func splitPrimes(numbers []int) (primesChan chan int, compositesChan chan int) {
	primesChan = make(chan int)
	compositesChan = make(chan int)

	go func() {
		defer close(primesChan)
		for _, n := range numbers {
			if primes.IsPrime(n) {
				primesChan <- n
			}
		}
	}()

	go func() {
		defer close(compositesChan)
		for _, n := range numbers {
			if !primes.IsPrime(n) {
				compositesChan <- n
			}
		}
	}()

	return primesChan, compositesChan
}

func goPrinter(ch <-chan int) chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)
		fmt.Println()
		for i := range ch {
			fmt.Println(i)
		}
	}()

	return done
}
