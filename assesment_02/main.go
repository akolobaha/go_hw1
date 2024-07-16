package main

import (
	"errors"
	"fmt"
)

func dfs(graph [][]int, visited []bool, start int) {
	fmt.Printf("%d ", start)
	visited[start] = true

	for i := 0; i < len(graph[start]); i++ {
		if graph[start][i] > 0 && !visited[i] {
			dfs(graph, visited, i)
		}
	}
}

func main() {
	mtx1 := [][]int{
		{0, 1, 3, 0, 0},
		{1, 0, 0, 1, 1},
		{3, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}

	fmt.Println(validate(mtx1, []int{4, 1, 0}))

	fmt.Println(calMaxGrade(mtx1))
}

func traverseGraph(matrix [][]int) {
	visited := make([]bool, len(matrix))
	stack := []int{0}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[node] {
			visited[node] = true
			//fmt.Println(node)

			for i := 0; i < len(matrix[node]); i++ {
				if matrix[node][i] > 0 && !visited[i] {
					stack = append(stack, i)
				} else {
					fmt.Printf("%d ", node)
				}
			}
		}
	}
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

func calMaxGrade(matrix [][]int) int {
	startVertex := 0
	visited := make([]bool, len(matrix))
	visited[startVertex] = true
	maxSum := 0

	findMaxPath(matrix, visited, startVertex, 0, &maxSum)

	return maxSum
}

func findMaxPath(graph [][]int, visited []bool, current int, sum int, maxSum *int) {
	if sum > *maxSum {
		*maxSum = sum
	}

	for i := 0; i < len(graph[current]); i++ {
		if graph[current][i] != 0 && !visited[i] {
			visited[i] = true
			findMaxPath(graph, visited, i, sum+graph[current][i], maxSum)
			visited[i] = false
		}
	}
}

// Реализовать стек
// Использовать стек для алгоритма поиска вглубину
// Найти как сохранить графа
// Максимальный балл

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

func calcUserGrade(matrix [][]int, userAnswer []int) int {
	return 0
}
