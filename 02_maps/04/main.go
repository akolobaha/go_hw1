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

func main() {
	printAggregatedResultTable()
}

/*
*
4. Для предыдущей задачи необходимо вывести сводную таблицу по всем предметам
*/

func printAggregatedResultTable() {
	data := readFile()
	unanimousResults := make(map[int]map[int][]float64)

	for _, result := range data.Results {
		if unanimousResults[result.ObjectId] == nil {
			unanimousResults[result.ObjectId] = make(map[int][]float64)
		}
		unanimousResults[result.ObjectId][Students[result.StudentId].Grade] = append(unanimousResults[result.ObjectId][Students[result.StudentId].Grade], float64(result.Result))
	}

	for objectId, objectResults := range unanimousResults {
		fmt.Printf("----------------------\n")
		fmt.Printf("%-12s | %-5s |\n", Objects[objectId].Name, "Mean")

		resultsByObject := make([]float64, 0)

		for grade, results := range objectResults {
			avgByGrade := avgGrade(results)
			resultsByObject = append(resultsByObject, results...)
			fmt.Printf("%-12d | %-5.1f |\n", grade, avgByGrade)
		}

		fmt.Printf("----------------------\n")
		fmt.Printf("%-12s | %-5.1f |\n", "mean", avgGrade(resultsByObject))
		fmt.Printf("----------------------\n")
	}
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

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func avgGrade(arr []float64) float64 {

	var result float64
	for _, value := range arr {
		result += value
	}

	if float64(len(arr)) != 0 {
		return result / float64(len(arr))
	}
	return 0
}
