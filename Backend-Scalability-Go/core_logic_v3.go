package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulation task structure
type Job struct {
	id int
}

// Worker function jo gprMax simulation ko simulate karega
func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d: gprMax simulation %d shuru ho rahi hai...\n", id, j.id)
		time.Sleep(time.Second * 2) // Simulation ka time
		fmt.Printf(
			"Worker %d: Simulation %d khatam!\n", id, j.id)
	}
}

func run() {
	const numJobs = 5
	jobs := make(chan Job, numJobs)
	var wg sync.WaitGroup

	// Sirf 3 workers chalenge (Parallelism limit)
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// Jobs queue mein daalna
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{id: j}
	}
	close(jobs)

	wg.Wait()
	fmt.Println("Saari gprMax simulations puri ho gayi! Halla bol! 🚀")
}