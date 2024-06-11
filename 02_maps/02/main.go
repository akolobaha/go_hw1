package main

import "fmt"

func main() {
	fmt.Println(countVotes([]string{"Ann", "Kate", "Peter", "Kate", "Ann", "Ann", "Helen"}))
}

/*
*
02
Напишите функцию подсчета каждого голоса за кандидата. Входной аргумент - массив с именами кандидатов.
Результативный - массив структуры Candidate, отсортированный по убыванию количества голосов.
*/
type candidates struct {
	name  string
	votes int
}

func countVotes(slice []string) []candidates {
	result := make([]candidates, 0)
	for _, candidate := range slice {
		if candidateExists(candidate, result) {
			for resultIndex, resultCandidate := range result {
				if resultCandidate.name == candidate {
					result[resultIndex].votes += 1
				}
			}
		} else {
			result = append(result, candidates{name: candidate, votes: 1})
		}
	}

	return result
}

func candidateExists(name string, candidates []candidates) bool {
	for _, candidate := range candidates {
		if candidate.name == name {
			return true
		}
	}
	return false
}
