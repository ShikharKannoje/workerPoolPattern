package main

import (
	"fmt"
	"time"
)

const (
	MAX_GOROUNTINES = 500
	MAX_JOBS        = 5000
)

func main() {

	fmt.Println("Enter the number of workers")
	var numWorkers int
	fmt.Scanln(&numWorkers)
	if numWorkers > MAX_GOROUNTINES {
		return
	}
	fmt.Println("Enter the number of jobs")
	var numJobs int
	fmt.Scanln(&numJobs)
	if numJobs > MAX_JOBS {
		return
	}

	jobsChan := make(chan int, numJobs)
	resultChan := make(chan int, numJobs)

	// starting the worker pool
	workerPool(numWorkers, jobsChan, resultChan)

	// sending the jobs
	for i := 0; i < numJobs; i++ {
		jobsChan <- i
	}
	close(jobsChan)

	// collecting results
	for i := 0; i < numJobs; i++ {
		resu := <-resultChan
		fmt.Println("result ", resu)
	}
	fmt.Println("Done")
}

func workerPool(numberOfWorker int, jobsChan <-chan int, resultChan chan<- int) {

	for i := 0; i < numberOfWorker; i++ {
		go worker(jobsChan, resultChan)
	}
}

func worker(jobsChan <-chan int, resultChan chan<- int) {
	for i := range jobsChan {
		resultChan <- process(i)
	}
}

func process(job int) int {
	time.Sleep(time.Second * 1)
	return job * 2
}
