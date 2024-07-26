package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const url = "localhost:8080"
const requests = 30

func main() {
	wg := &sync.WaitGroup{}
	client(wg)

}

func client(wg *sync.WaitGroup) {
	wg.Add(requests)
	for range requests {
		go func() {
			defer wg.Done()

			resp, _ := http.Get("http://" + url)

			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(resp.StatusCode, string(body))
		}()
	}

	wg.Wait()
}
