package main

import "fmt"

func main() {
	chan1 := counter(5)
	chan2 := counter(5)

	united := uniteChannels(chan1, chan2)

	done := goPrinter(united)
	<-done
}

func counter(num int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= num; i++ {
			ch <- i
		}
	}()

	return ch
}

func uniteChannels(ch1 <-chan int, ch2 <-chan int) chan int {
	result := make(chan int)

	go func() {
		defer close(result)
		for n := range ch1 {
			result <- n
		}
		for n := range ch2 {
			result <- n
		}
	}()

	return result
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
