package main

import (
	"fmt"
	"sync"
)

func workerJob(index int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", index, j)
		results <- j
	}
}

func main() {
	const jobsCount = 5
	jobs := make(chan int, jobsCount)
	results := make(chan int, jobsCount)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go workerJob(w, jobs, results, &wg)
	}

	for j := 1; j <= jobsCount; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result)
	}

}
