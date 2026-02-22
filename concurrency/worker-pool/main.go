package main

import (
	"fmt"
	"time"
)

func main() {
	numJobs := 7
	numWorkers := 3

	jobs := make(chan string, numJobs)
	results := make(chan string, numJobs)

	for i := range numWorkers {
		go worker(i, jobs, results)
	}

	todos := []string{
		"push-ups",
		"pull-ups",
		"muscle-ups",
		"lunges",
		"crunch",
		"squat",
		"burpee",
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Start to assign jobs!")

	for _, todo := range todos {
		jobs <- todo
	}

	close(jobs)

	for range numJobs {
		fmt.Println(<-results)
	}
}

func worker(id int, jobs <-chan string, results chan string) {
	fmt.Println("worker", id, "is ready to work!")

	for job := range jobs {
		fmt.Println("worker", id, "receives the job:", job)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished the job:", job)
		results <- fmt.Sprintf("%s is done", job)
	}
}
