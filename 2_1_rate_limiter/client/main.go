package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var url = "localhost:8080"
var requests = 30

func main() {
	//go server()
	wg := &sync.WaitGroup{}
	client(wg)

}

func client(wg *sync.WaitGroup) {
	for range requests {
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
