package main

func main() {

	arr1 := []int{1, 2, 3, 4, 5, 4, 222}
	arr2 := []int{4, 1, 2, 3, 5, 5, 222}

	println(isEqualArrays(arr1, arr2))
}

func isEqualArrays[T comparable](arr1 []T, arr2 []T) bool {
	resultingMap := make(map[T]bool)

	for i := range arr1 {
		resultingMap[arr1[i]] = false
	}

	for i := range arr2 {
		_, ok := resultingMap[arr2[i]]
		if !ok {
			return false
		} else {
			resultingMap[arr2[i]] = true
		}
	}

	for _, val := range resultingMap {
		if !val {
			return false
		}
	}

	return true
}
