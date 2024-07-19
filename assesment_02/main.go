package main

import (
	"errors"
	"fmt"
)

func main() {

	mtx0 := [][]int{
		{0, 2, 3, 0, 0, 0},
		{2, 0, 0, 1, 1, 0},
		{3, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 7},
		{0, 0, 0, 0, 7, 0},
	}

	fmt.Println(calMaxGrade(mtx0))
	fmt.Println(calcUserGrade(mtx0, []int{4, 5}))
}

func calcUserGrade(matrix [][]int, userAnswer []int) (int, error) {
	_, err := validateBoth(matrix, userAnswer)
	if err != nil {
		return 0, err
	}

	sum := 0
	for i := 0; i < len(userAnswer)-1; i++ {
		from := userAnswer[i]
		to := userAnswer[i+1]
		sum += matrix[from][to]
	}
	return sum, nil
}

func calMaxGrade(matrix [][]int) (int, error) {
	_, err := validateMatrix(matrix)
	if err != nil {
		return 0, err
	}

	result := 0

	for startVertex := 0; startVertex < len(matrix); startVertex++ {
		maxSum := 0
		visited := make([]bool, len(matrix))
		visited[startVertex] = true
		calculatePath(matrix, visited, startVertex, 0, &maxSum)
		if maxSum > result {
			result = maxSum
		}
	}

	return result, nil
}

func calculatePath(graph [][]int, visited []bool, current int, sum int, maxSum *int) {
	if sum > *maxSum {
		*maxSum = sum
	}

	for i := 0; i < len(graph[current]); i++ {
		if graph[current][i] != 0 && !visited[i] {
			visited[i] = true
			calculatePath(graph, visited, i, sum+graph[current][i], maxSum)
			visited[i] = false
		}
	}
}

func validateMatrix(matrix [][]int) (bool, error) {
	err := error(nil)
	_, err = matrixIsSquare(matrix)
	if err != nil {
		return false, err
	}

	_, err = matrixNoLoop(matrix)
	if err != nil {
		return false, err
	}

	return true, nil
}

func validateBoth(matrix [][]int, slice []int) (bool, error) {
	err := error(nil)
	_, err = answerLength(matrix, slice)
	if err != nil {
		return false, err
	}
	_, err = answerIsUnique(slice)
	if err != nil {
		return false, err
	}
	_, err = matrixIsSquare(matrix)
	if err != nil {
		return false, err
	}
	_, err = matrixNoLoop(matrix)
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

func EvalSequence(matrix [][]int, userAnswer []int) (int, error) {
	// validation
	maxGrade, err := calMaxGrade(matrix)

	if err != nil {
		return 0, err
	}

	userGrade, err := calcUserGrade(matrix, userAnswer)

	if err != nil {
		return 0, err
	}

	percent := userGrade * 100 / maxGrade

	return percent, nil
}
