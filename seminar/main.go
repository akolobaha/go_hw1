package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	chanCubes := calcCubes(nums)
	chanSquares := calcSquares(nums)

	var cubes = make([]int, len(chanCubes))
	var squares = make([]int, len(chanSquares))

	for v := range chanCubes {
		cubes = append(cubes, v)
	}

	for v := range chanSquares {
		squares = append(squares, v)
	}

	fmt.Println(cubes, squares)
}

func calcCubes(nums []int) chan int {
	channel := make(chan int)
	go func() {
		defer close(channel)
		for num := range nums {
			channel <- num * num * num
		}
	}()

	return channel
}

func calcSquares(nums []int) chan int {
	channel := make(chan int)
	go func() {
		defer close(channel)
		for num := range nums {
			channel <- num * num
		}
	}()
	return channel
}
