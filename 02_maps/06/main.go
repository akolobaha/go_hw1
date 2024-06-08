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

func main() {
	printResultsTable()
}

func printResultsTable() {
	data := readFile()

	fmt.Printf("--------------------------------------------\n")
	fmt.Printf("%-12s | %-5s | %-10s | %-6s |\n", "Student name", "Grade", "Object", "Result")
	fmt.Printf("--------------------------------------------\n")

	for _, result := range data.Results {
		student := FilterOne(data.Students, result.StudentId, getStudentById)
		object := FilterOne(data.Objects, result.ObjectId, getObjectNameById)

		//student := getStudentById(result.StudentId, data.Students)
		fmt.Printf("%-12s | %-5d | %-10s | %-6d |\n", student.Name, student.Grade, object.Name, result.Result)
	}
}

func FilterOne[T any](s []T, id int, f func(T, int) bool) T {
	var r T
	for _, v := range s {
		if f(v, id) {
			return v
		}
	}
	return r
}

func getObjectNameById(object Object, id int) bool {
	return object.Id == id
}

func getStudentById(student Student, id int) bool {
	return student.Id == id
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
	if err != nil {
		log.Fatal(err)
	}

	return data
}
