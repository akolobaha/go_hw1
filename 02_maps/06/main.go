package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

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

var Students = make(map[int]Student)
var Objects = make(map[int]Object)
var Grades = make(map[int]struct{})

/*
 6. Перепишите задачу #4 с использованием функций высшего порядка, изученных на лекции.

Желательно реализуйте эти функции самостоятельно.
*/
func main() {
	printAggregatedResultTable()
}

/*
*
4. Для предыдущей задачи необходимо вывести сводную таблицу по всем предметам
*/

func printAggregatedResultTable() {
	data := readFile()

	for _, object := range data.Objects {
		resultsGroupedByObjects := Filter(data.Results, object.Id, filterResultByObjectId)

		fmt.Printf("----------------------\n")
		fmt.Printf("%-12s | %-5s |\n", Objects[object.Id].Name, "Mean")

		for grade, _ := range Grades {
			resultGroupedByGrade := Filter(resultsGroupedByObjects, grade, filterStudentByGrade)

			if len(resultGroupedByGrade) > 0 {
				mappedResults := Map(resultGroupedByGrade, mapResults)
				resultsByObject := make([]int, 0)
				avgByGrade := avgGrade(mappedResults)
				resultsByObject = append(resultsByObject, mappedResults...)
				fmt.Printf("%-12d | %-5.1f |\n", grade, avgByGrade)
			}

		}

		mean := avgGrade(Map(resultsGroupedByObjects, mapResults))

		fmt.Printf("----------------------\n")
		fmt.Printf("%-12s | %-5.1f |\n", "mean", mean)
		fmt.Printf("----------------------\n")
	}
}

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func mapResults(result Result) int {
	return result.Result
}

func Filter[T any](s []T, id int, f func(T, int) bool) []T {
	var r []T
	for _, v := range s {
		if f(v, id) {
			r = append(r, v)
		}
	}
	return r
}

func filterResultByObjectId(result Result, objectId int) bool {
	return result.ObjectId == objectId
}

func filterStudentByGrade(result Result, grade int) bool {
	return Students[result.StudentId].Grade == grade
}

func readFile() Data {
	file, err := os.Open("../dz3.json")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		os.Exit(1)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data Data
	err = decoder.Decode(&data)

	for _, v := range data.Students {
		Students[v.Id] = v
	}

	for _, v := range data.Objects {
		Objects[v.Id] = v
	}

	for _, value := range Students {
		Grades[value.Grade] = struct{}{}
	}

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func avgGrade(arr []int) float64 {

	var result int
	for _, value := range arr {
		result += value
	}

	if float64(len(arr)) != 0 {
		return float64(result) / float64(len(arr))
	}
	return 0
}
