package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 10

	oddTurn := make(chan struct{})
	evenTurn := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	// Odd Number Goroutine
	go func() {
		defer wg.Done()

		for i := 1; i <= n; i += 2 {
			<-oddTurn

			fmt.Println(i)

			if i+1 <= n {
				evenTurn <- struct{}{}
			}
		}
	}()

	// Even Number Goroutine
	go func() {
		defer wg.Done()

		for i := 2; i <= n; i += 2 {
			<-evenTurn

			fmt.Println(i)

			if i+1 <= n {
				oddTurn <- struct{}{}
			}
		}
	}()

	// Start execution with odd goroutine
	oddTurn <- struct{}{}

	wg.Wait()
}
