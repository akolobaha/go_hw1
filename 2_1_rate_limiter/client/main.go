package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var url = "localhost:8080"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	//go server()
	wg := &sync.WaitGroup{}
	client(wg)

}

//func server() {
//	http.HandleFunc("/", handler)
//
//	fmt.Println("Сервер запущен на http://localhost:8080/")
//	if err := http.ListenAndServe(url, nil); err != nil {
//		fmt.Println("Ошибка при запуске сервера:", err)
//	}
//}

func client(wg *sync.WaitGroup) {
	for range 10 {
		go func() {
			defer wg.Done()

			wg.Add(1)
			resp, err := http.Get("http://" + url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
			}
			fmt.Println(resp.StatusCode, string(body))
		}()
	}

	wg.Wait()
}
