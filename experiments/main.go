package main

import (
	"fmt"
	"sync"
)

type Job struct {
	id int
}

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.id)
	}
}

func main() {
	numWorkers := 3
	numJobs := 5

	jobs := make(chan Job, numJobs)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- Job{id: i}
	}
	close(jobs)

	wg.Wait()
}
