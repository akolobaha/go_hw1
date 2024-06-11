package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(crossingSlices([]int{1, 2, 3, 2}, []int{3, 2}, []int{}))
}

/*
01
Напишите функцию, которая находит пересечение неопределенного количества слайсов типа int.
Каждый элемент в пересечении должен быть уникальным. Слайс-результат должен быть отсортирован в восходящем порядке.
*/
func crossingSlices(slices ...[]int) []int {
	shortestSlide := getShortestSlide(slices...)
	result := make([]int, 0)

outerLoop:
	for _, elem := range shortestSlide {
		for _, slice := range slices {
			if !sliceContainsElem(slice, elem) {
				continue outerLoop
			}
		}

		if !sliceContainsElem(result, elem) {
			result = append(result, elem)
		}
	}
	sort.Ints(result)
	return result
}

func getShortestSlide(slices ...[]int) []int {
	var minSliceIndex int

	for slice := range slices {
		if len(slices[slice]) < len(slices[minSliceIndex]) {
			minSliceIndex = slice
		}
	}
	return slices[minSliceIndex]
}

func sliceContainsElem(slice []int, elem int) bool {
	for _, item := range slice {
		if item == elem {
			return true
		}
	}
	return false
}
