package main

import "fmt"

type Numbers []int

func main() {
	arr := Numbers{1, 2, 3, 4, 5}
	arr2 := Numbers{1, 2, 3, 4, 5, 6}

	fmt.Println(arr.Sum())
	fmt.Println(arr.Mul())
	fmt.Println(arr.hasElement(3))

	fmt.Println(arr.isEqualArrays(arr2))

	newArr1 := arr.removeElementByIndex(3)
	fmt.Println(newArr1)
	newArr2 := newArr1.removeElementByValue(5)
	fmt.Println(newArr2)
}

func (numbers Numbers) Sum() (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func (numbers Numbers) Mul() (res int) {
	res = 1
	for _, number := range numbers {
		res *= number
	}
	return res
}

func (numbers Numbers) isEqualArrays(toCompare Numbers) bool {
	resultingMap := make(map[int]bool)

	for i := range numbers {
		resultingMap[numbers[i]] = false
	}

	for i := range toCompare {
		_, ok := resultingMap[toCompare[i]]
		if !ok {
			return false
		} else {
			resultingMap[toCompare[i]] = true
		}
	}

	for _, val := range resultingMap {
		if !val {
			return false
		}
	}

	return true
}

func (numbers Numbers) hasElement(arg int) int {
	for i, number := range numbers {
		if number == arg {
			return i
		}
	}
	return -1
}

func (numbers Numbers) removeElementByIndex(index int) Numbers {
	return append(numbers[:index], numbers[index+1:]...)
}

func (numbers Numbers) removeElementByValue(arg int) Numbers {
	for i, number := range numbers {
		if number == arg {
			return numbers.removeElementByIndex(i)
		}
	}
	return numbers
}
