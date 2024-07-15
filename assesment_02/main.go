package main

import (
	"errors"
	"fmt"
)

func main() {
	mtx1 := [][]int{
		{0, 2, 3, 0, 0},
		{2, 0, 0, 1, 1},
		{3, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}

	fmt.Println(validate(mtx1, []int{4, 1, 0}))
}

func validate(matrix [][]int, slice []int) (bool, error) {

	_, err := answerLength(matrix, slice)
	if err != nil {
		return false, err
	}
	_, err = answerIsUnique(slice)
	if err != nil {
		return false, err
	}
	_, err = matrixNoLoop(matrix)
	if err != nil {
		return false, err
	}
	_, err = matrixIsSquare(matrix)
	if err != nil {
		return false, err
	}

	return true, nil
}

func answerLength(matrix [][]int, slice []int) (bool, error) {
	if len(matrix) >= len(slice) {
		return true, nil
	}
	return false, errors.New("matrix length does not match slice length")
}

func answerIsUnique(slice []int) (bool, error) {
	seen := make(map[int]bool)
	for _, val := range slice {
		if seen[val] {
			return false, errors.New("duplicate answer element")
		}
		seen[val] = true
	}
	return true, nil
}

func matrixNoLoop(matrix [][]int) (bool, error) {
	for i := 0; i < len(matrix); i++ {
		if matrix[i][i] != 0 {
			return false, errors.New("matrix has loop")
		}
	}
	return true, nil
}

func matrixIsSquare(matrix [][]int) (bool, error) {
	for i := 0; i < len(matrix); i++ {
		if len(matrix) != len(matrix[0]) {
			return false, errors.New("matrix is not square shape")
		}
	}
	return true, nil
}

// Найти как сохранить графа

// Получили переданную матрицы, записали ее в память виде графа
// Делаем обход графа
// Сопостовляем стоимости ребер графа с переданным вариантом ответа

func EvalSequence(matrix [][]int, userAnswer []int) int {

	// validation
	maxGrade := calMaxGrade(matrix)
	userGrade := calcUserGrade(matrix, userAnswer)

	percent := userGrade * 100 / maxGrade

	return percent
}

func calMaxGrade(matrix [][]int) int {
	return -1
}

func calcUserGrade(matrix [][]int, userAnswer []int) int {
	return 0
}
