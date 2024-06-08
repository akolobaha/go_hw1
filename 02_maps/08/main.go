package main

import "fmt"

func Len[T comparable](value []T) int {
	return len(value)
}

// дописать метод compare и swap для comparable??

//func isEqualArrays[T myComparable](arr1, arr2 []T) bool {
//	sort.Sort(T)
//}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 9}
	//arr3 := []int{1, 3, 4, 5, 6, 7, 8, 10, 9}

	fmt.Println(Len(arr1))
}
