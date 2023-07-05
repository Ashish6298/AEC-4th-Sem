package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numPhilosophers = 3
	maxDiningCycles = 2
)

type Philospher struct {
	id                  int
	leftFork, rightFork *sync.Mutex
	diningCycles        int
}

type DiningTable struct {
	philospher []*Philospher
	waiter     *sync.Mutex
}

func (p *Philospher) think() {
	fmt.Printf("Philospher %d is thinking\n", p.id)
	time.Sleep(time.Second)
}

func (p *Philospher) eat() {
	p.leftFork.Lock()
	p.rightFork.Lock()
	fmt.Printf("Philospher %d is eating\n", p.id)
	time.Sleep(time.Second)
	p.rightFork.Unlock()
	p.leftFork.Unlock()
	p.diningCycles++
}

func (p *Philospher) dine(table *DiningTable) {
	for p.diningCycles < maxDiningCycles {
		p.think()
		table.waiter.Lock()
		p.eat()
		table.waiter.Unlock()
	}
}

func main() {

	table := &DiningTable{
		philospher: make([]*Philospher, numPhilosophers),
		waiter:     &sync.Mutex{},
	}
	// create forks
	fork := make([]*sync.Mutex, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		fork[i] = &sync.Mutex{}
	}

	// create philosophers and assign forks
	for i := 0; i < numPhilosophers; i++ {
		table.philospher[i] = &Philospher{
			id:        i,
			leftFork:  fork[i],
			rightFork: fork[(i+1)%numPhilosophers],
		}
	}

	// start dining
	var wg sync.WaitGroup
	wg.Add(numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		go func(p *Philospher) {
			defer wg.Done()
			p.dine(table)
		}(table.philospher[i])
	}
	wg.Wait()
}


/*
		Output:
		Philospher 2 is thinking
		Philospher 0 is thinking
		Philospher 1 is thinking
		Philospher 1 is eating
		Philospher 1 is thinking
		Philospher 2 is eating
		Philospher 2 is thinking
		Philospher 1 is eating
		Philospher 0 is eating
		Philospher 0 is thinking
		Philospher 2 is eating
		Philospher 0 is eating
*/