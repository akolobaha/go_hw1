package main

func main() {

	arr1 := []int{1, 2, 3, 4, 5, 4, 222}
	arr2 := []int{4, 1, 2, 3, 5, 5, 222}

	println(isEqualArrays(arr1, arr2))
}

func isEqualArrays[T comparable](arr1 []T, arr2 []T) bool {
	return arrInArr(arr1, arr2) && arrInArr(arr2, arr1)
}

func arrInArr[T comparable](arr1 []T, arr2 []T) bool {
outer:
	for i := range arr1 {
		for j := range arr2 {
			if arr1[i] == arr2[j] {
				continue outer
			}
		}
		return false
	}
	return true
}
