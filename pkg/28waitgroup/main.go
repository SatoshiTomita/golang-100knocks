package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numWorkers := 5

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			fmt.Printf("Worker %d is working\n", workerID)
		}(i)
	}

	wg.Wait()
	fmt.Println("All workers have finished")
}
