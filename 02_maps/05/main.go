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

type Cache[K comparable, V any] struct{ m map[K]V }

var CacheStudent Cache[int, Student]
var CacheObject Cache[int, Object]

func (c *Cache[K, V]) Init() {
	c.m = make(map[K]V)
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.m[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	k, ok := c.m[key]
	return k, ok
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
		student, _ := CacheStudent.Get(result.StudentId)
		object, _ := CacheObject.Get(result.ObjectId)
		fmt.Printf("%-12s | %-5d | %-10s | %-6d |\n", student.Name, student.Grade, object.Name, result.Result)
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
	if err != nil {
		log.Fatal(err)
	}

	CacheStudent.Init()
	CacheObject.Init()

	for _, student := range data.Students {
		CacheStudent.Set(student.Id, student)
	}

	for _, object := range data.Objects {
		CacheObject.Set(object.Id, Object{
			Id:   object.Id,
			Name: object.Name,
		})
	}

	return data
}
