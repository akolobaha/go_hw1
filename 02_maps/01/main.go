package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(crossingSlices([]int{1, 2, 3, 2, 7}, []int{3, 2, 7}, []int{1, 3, 7, 7, 7}))
}

/*
01
Напишите функцию, которая находит пересечение неопределенного количества слайсов типа int.
Каждый элемент в пересечении должен быть уникальным. Слайс-результат должен быть отсортирован в восходящем порядке.
*/
func crossingSlices(slices ...[]int) []int {
	result := make([]int, 0)
	resultMap := make(map[int]int)

	for _, slice := range slices {
		unique := make(map[int]struct{})
		for _, value := range slice {
			unique[value] = struct{}{}
		}

		for key, _ := range unique {
			resultMap[key]++
		}
	}

	for key := range resultMap {
		if resultMap[key] == len(slices) {
			result = append(result, key)
		}
	}

	sort.Ints(result)
	return result
}
