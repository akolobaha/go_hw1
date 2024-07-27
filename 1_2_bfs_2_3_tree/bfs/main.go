package main

import (
	"fmt"
)

func bfs(graph [][]int, start int) []int {
	visited := make([]bool, len(graph))
	queue := []int{start}
	visited[start] = true
	result := []int{}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		result = append(result, vertex)

		// Проверка соседей
		for i := 0; i < len(graph[vertex]); i++ {
			if graph[vertex][i] != 0 && !visited[i] {
				visited[i] = true
				queue = append(queue, i)
			}
		}
	}

	return result
}

func main() {
	mtx0 := [][]int{
		{0, 2, 3, 0, 0, 0},
		{2, 0, 0, 1, 1, 0},
		{3, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 7},
		{0, 0, 0, 0, 7, 0},
	}

	fmt.Println("BFS обход:", bfs(mtx0, 1))
	fmt.Println("BFS обход:", bfs(mtx0, 2))
	fmt.Println("BFS обход:", bfs(mtx0, 3))
}
