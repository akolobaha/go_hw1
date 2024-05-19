package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	fileName := flag.String("f", "problems.csv", "Имя csv файла для обработки")
	random := flag.Bool("r", false, "Чтение строк в произвольном порядке")

	flag.Parse()

	lines, _ := readFile(*fileName)
	var correct, wrong int = processFile(lines, *random)

	println("Количество правильных ответов: ", correct)
	println("Количество не верных ответов: ", wrong)

}

func readFile(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func processFile(lines [][]string, random bool) (correctAnswer, wrongAnswers int) {
	correctAnswers := 0

	if random {
		rand.Shuffle(len(lines), func(i, j int) { lines[i], lines[j] = lines[j], lines[i] })
	}

	for _, line := range lines {
		question, result := parseLine(line)
		guess := getUserInput(question)

		if answerIsCorrect(result, guess) {
			correctAnswers++
		}
	}

	return correctAnswers, len(lines) - correctAnswers
}

func parseLine(line []string) (question, result string) {
	return line[0], line[1]
}

func getUserInput(question string) string {
	var guess string
	fmt.Println(question)
	fmt.Scanln(&guess)
	return guess
}

func answerIsCorrect(result, guess string) bool {
	guess = strings.ToLower(strings.TrimSpace(guess))
	result = strings.ToLower(strings.TrimSpace(result))
	return result == guess
}
