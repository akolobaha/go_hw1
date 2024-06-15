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

var data Data

func main() {
	printResultsTable()

}

func printResultsTable() {
	data = readFile()

	fmt.Printf("--------------------------------------------\n")
	fmt.Printf("%-12s | %-5s | %-10s | %-6s |\n", "Student name", "Grade", "Object", "Result")
	fmt.Printf("--------------------------------------------\n")

	var excelentStudents = Filter(data.Students, isExcelentStudent)

	for _, result := range data.Results {
		student := FilterOne(data.Students, result.StudentId, getStudentById)
		if !inArray(student, excelentStudents) {
			continue
		}

		object := FilterOne(data.Objects, result.ObjectId, getObjectNameById)

		//student := getStudentById(result.StudentId, data.Students)
		fmt.Printf("%-12s | %-5d | %-10s | %-6d |\n", student.Name, student.Grade, object.Name, result.Result)
	}
}

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {

		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func isExcelentStudent(student Student) bool {
	var isExcelent bool = true
	for _, result := range data.Results {
		if result.StudentId == student.Id && result.Result != 5 {
			isExcelent = false
			break
		}
	}

	return isExcelent
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

func inArray[T comparable](student T, students []T) bool {
	for _, result := range students {
		if student == result {
			return true
		}
	}
	return false
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
