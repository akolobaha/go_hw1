package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	fmt.Println(crossingSlices([]int{1, 2, 3, 2}, []int{3, 2}, []int{}))
	fmt.Println(countVotes([]string{"Ann", "Kate", "Peter", "Kate", "Ann", "Ann", "Helen"}))

	printTable()
}

/*
01
Напишите функцию, которая находит пересечение неопределенного количества слайсов типа int.
Каждый элемент в пересечении должен быть уникальным. Слайс-результат должен быть отсортирован в восходящем порядке.
*/
func crossingSlices(slices ...[]int) []int {
	shortestSlide := getShortestSlide(slices...)
	result := make([]int, 0)

outerLoop:
	for _, elem := range shortestSlide {
		for _, slice := range slices {
			if !sliceContainsElem(slice, elem) {
				continue outerLoop
			}
		}

		if !sliceContainsElem(result, elem) {
			result = append(result, elem)
		}
	}
	sort.Ints(result)
	return result
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

/*
*
03
У учеников старших классов прошел контрольный срез по нескольким предметам. Выведите данные в читаемом виде
в таблицу вида
*/

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

type Object struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Result struct {
	ObjectId  int `json:"object_id"`
	StudentId int `json:"student_id"`
	Result    int `json:"result"`
}

type Data struct {
	Students []Student `json:"students"`
	Objects  []Object  `json:"objects"`
	Results  []Result  `json:"results"`
}

func printTable() {
	file, err := os.Open("dz3.json")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data Data
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("--------------------------------------------\n")
	fmt.Printf("%-12s | %-5s | %-10s | %-6s |\n", "Student name", "Grade", "Object", "Result")
	fmt.Printf("--------------------------------------------\n")

	for _, result := range data.Results {
		student := getStudentById(result.StudentId, data.Students)
		fmt.Printf("%-12s | %-5d | %-10s | %-6d |\n", student.Name, student.Grade, getObjectNameById(result.ObjectId, data.Objects), result.Result)
	}
}

func getObjectNameById(id int, objects []Object) string {
	for _, object := range objects {
		if object.Id == id {
			return object.Name
		}
	}
	return ""
}

func getStudentById(id int, students []Student) Student {
	for _, student := range students {
		if student.Id == id {
			return student
		}
	}
	return Student{}
}

/*
*
Вспомогательные функции
*/

func getShortestSlide(slices ...[]int) []int {
	var minSliceIndex int

	for slice := range slices {
		if len(slices[slice]) < len(slices[minSliceIndex]) {
			minSliceIndex = slice
		}
	}
	return slices[minSliceIndex]
}

func sliceContainsElem(slice []int, elem int) bool {
	for _, item := range slice {
		if item == elem {
			return true
		}
	}
	return false
}

func candidateExists(name string, candidates []candidates) bool {
	for _, candidate := range candidates {
		if candidate.name == name {
			return true
		}
	}
	return false
}
