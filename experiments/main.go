package main

import "fmt"

var mtx2 = [][]int{
	{0, 2, 3, 0, 0},
	{2, 0, 0, 1, 1},
	{3, 0, 0, 0, 0},
	{0, 1, 0, 0, 0},
	{0, 1, 0, 0, 0},
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

func main() {

	maxSum := calMaxGrade(mtx2)

	fmt.Println("Самый дорогой путь в графе:", maxSum)
}
