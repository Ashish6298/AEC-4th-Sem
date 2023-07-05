package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, checkpoint chan bool, resume chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Duration(id) * time.Second) // simulate work

	fmt.Printf("Worker %d waiting for checkpoint\n", id)
	checkpoint <- true // signal that checkpoint has been reached
	<-resume           // wait for permission signal
	fmt.Printf("Worker %d resuming\n", id)
	// continue with the remaining work
}

func main() {
	/*
		Output:
		Worker 4 starting
		Worker 0 starting
		Worker 0 waiting for checkpoint
		Worker 1 starting
		Worker 2 starting
		Worker 3 starting
		Worker 1 waiting for checkpoint
		Worker 2 waiting for checkpoint
		Worker 3 waiting for checkpoint
		Worker 4 waiting for checkpoint
		All workers reached the checkpoint
		Resuming all workers
		Worker 1 resuming
		Worker 4 resuming
		Worker 2 resuming
		Worker 3 resuming
		Worker 0 resuming
		All workers completed their work
	*/
	numWorkers := 5
	checkpoint := make(chan bool)
	resume := make(chan bool)
	var wg sync.WaitGroup

	// launch the workers goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, checkpoint, resume, &wg)
	}

	// wait for all workers to reach the checkpoint
	for i := 0; i < numWorkers; i++ {
		<-checkpoint
	}

	fmt.Println("All workers reached the checkpoint")

	fmt.Println("Resuming all workers")

	// signal all workers to resume
	for i := 0; i < numWorkers; i++ {
		resume <- true
	}
	wg.Wait()
	fmt.Println("All workers completed their work")
}