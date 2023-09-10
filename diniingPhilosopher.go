package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	start := make(chan struct{})
	var num int
	fmt.Printf("Enter the number of pjilosopher")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			<-start

			fmt.Printf("Philosopher %d is thinking\n", id)
			time.Sleep(time.Second)

			fmt.Printf("Philosopher %d is started to eat\n", id)
			time.Sleep(time.Second)
			time.Sleep(time.Second)

			fmt.Printf("Philosopher %d is finished  eating\n", id)
			time.Sleep(time.Second)
		}(i)
	}

	fmt.Println(" Philosopher Starting  their tasks....")
	close(start)
	wg.Wait()
	fmt.Println("All philosopher completed their task")
}
