package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int)

	//producer
	go producer(jobs)

	consumer(jobs)
}

func producer(jobs chan<- int) {
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
}

func consumer(jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Consumed: %d\n", job)
	}
}
