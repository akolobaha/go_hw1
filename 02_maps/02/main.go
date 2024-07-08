package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countVotes([]string{"Ann", "Kate", "Peter", "Kate", "Ann", "Ann", "Helen"}))
}

/*
*
02
Напишите функцию подсчета каждого голоса за кандидата. Входной аргумент - массив с именами кандидатов.
Результативный - массив структуры Candidate, отсортированный по убыванию количества голосов.
*/

type Candidate struct {
	name  string
	votes int
}

type Candidates []Candidate

func (a Candidates) Len() int           { return len(a) }
func (a Candidates) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Candidates) Less(i, j int) bool { return a[i].votes > a[j].votes }

func countVotes(slice []string) Candidates {
	mapResult := make(map[string]int)
	result := make(Candidates, 0)

	for _, candidate := range slice {
		mapResult[candidate]++
	}

	for name, votes := range mapResult {
		result = append(result, Candidate{name: name, votes: votes})
	}

	sort.Sort(result)

	return result
}
