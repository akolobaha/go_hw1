package main

import "fmt"

func main() {
	//slice1 := make([]int, 1, 2, 3, 2)
	//slice2 := make([]int, 3, 5)

	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	//slice3 := []int{7, 8, 9}

	fmt.Println(slice1, slice2)

	crossingSlices(slice1, slice2)

	fmt.Println("_main()")
}

/*
01
Напишите функцию, которая находит пересечение неопределенного количества слайсов типа int.
Каждый элемент в пересечении должен быть уникальным. Слайс-результат должен быть отсортирован в восходящем порядке.
*/
func crossingSlices(slices ...[]int) {
	for slice := range slices {
		fmt.Println(slice)
	}
}
