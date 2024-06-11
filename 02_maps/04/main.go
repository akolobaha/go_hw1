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
	printAggregatedResultTable()
}

/*
*
4. Для предыдущей задачи необходимо вывести сводную таблицу по всем предметам
*/

type subjectGrade struct {
	grade  int
	result []int
}

func printAggregatedResultTable() {
	data := readFile()

	subjectsResult := make(map[string]map[int]subjectGrade)

	for _, result := range data.Results {
		student := getStudentById(result.StudentId, data.Students)
		objectName := getObjectNameById(result.ObjectId, data.Objects)

		value, ok := subjectsResult[objectName]
		if !ok {
			value = make(map[int]subjectGrade)
			subjectsResult[objectName] = value
		}

		subjectGradeValue, ok := value[student.Grade]
		if ok {
			subjectGradeValue.result = append(subjectGradeValue.result, result.Result)
		} else {
			subjectGradeValue = subjectGrade{
				grade:  student.Grade,
				result: []int{result.Result},
			}
		}

		value[student.Grade] = subjectGradeValue
	}

	for subjectName, subjectData := range subjectsResult {
		fmt.Printf("----------------------\n")
		fmt.Printf("%-12s | %-5s |\n", subjectName, "Mean")
		var totalObjGrade float64
		for gradeNumber, gradeData := range subjectData {
			avgGrade := avgGrade(gradeData.result)
			totalObjGrade += avgGrade
			fmt.Printf("%-12d | %-5.1f |\n", gradeNumber, avgGrade)
		}
		fmt.Printf("----------------------\n")
		fmt.Printf("%-12s | %-5.1f |\n", "mean", float64(totalObjGrade)/float64(len(subjectData)))
		fmt.Printf("----------------------\n")
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

func avgGrade(arr []int) float64 {
	var result int
	for _, value := range arr {
		result += value
	}
	return float64(result) / float64(len(arr))
}
