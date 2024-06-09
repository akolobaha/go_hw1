package main

type Numbers []int

func main() {

}

func (numbers Numbers) Sum() (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func (numbers Numbers) Mul(res int) {
	for _, number := range numbers {
		res *= number
	}
	return
}

func (numbers Numbers) sliceLenIsEqual(toCompare []Numbers) bool {
	return len(numbers) == len(toCompare)
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
	slice := make(int, 40)
	return append(numbers[:index], numbers[index:])
}
