//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {

//}
//

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// DownloadFile скачивает файл по указанному URL и сохраняет его в заданной директории
func DownloadFile(url string, filepath string) error {
	// Отправляем GET-запрос
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer response.Body.Close()

	// Проверяем код статуса ответа
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка при скачивании файла: %s", response.Status)
	}

	// Создаем файл для сохранения
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %v", err)
	}
	defer out.Close()

	// Копируем данные из ответа в файл
	_, err = io.Copy(out, response.Body)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл: %v", err)
	}

	return nil
}

func worker(url string, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", url, "started job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(string(rune(w)), jobs, results)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 9; a++ {
		<-results
	}
	//url := []string{
	//	"https://www.catseyepest.com/wp-content/uploads/2021/10/iStock_000043526268_Large-1024x683.jpg",
	//} // Укажите URL файла
	//filepath := "./images/1.jpg" // Укажите путь для сохранения файла
	//
	//err := DownloadFile(url, filepath)
	//if err != nil {
	//	fmt.Println("Ошибка:", err)
	//} else {
	//	fmt.Println("Файл успешно скачан:", filepath)
	//}
}
