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

	printResultsTable()
}

/*
*
03
У учеников старших классов прошел контрольный срез по нескольким предметам. Выведите данные в читаемом виде
в таблицу вида
*/

func printResultsTable() {
	data := readFile()

	fmt.Printf("--------------------------------------------\n")
	fmt.Printf("%-12s | %-5s | %-10s | %-6s |\n", "Student name", "Grade", "Object", "Result")
	fmt.Printf("--------------------------------------------\n")

	for _, result := range data.Results {
		fmt.Printf("%-12s | %-5d | %-10s | %-6d |\n", Students[result.StudentId].Name, Students[result.StudentId].Grade, Objects[result.ObjectId].Name, result.Result)
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
