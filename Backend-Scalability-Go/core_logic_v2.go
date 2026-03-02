package main

import (
	"fmt"
	"time"
)

// Simulation struct jo task ka data hold karega
type SimTask struct {
	ID    int
	Name  string
}

// Function jo simulation ko "run" karega
func runSimulation(task SimTask) {
	fmt.Printf("🚀 Starting Physics Simulation: %s (ID: %d)\n", task.Name, task.ID)
	time.Sleep(2 * time.Second) // Fake processing time
	fmt.Printf("✅ Simulation %d Completed successfully!\n", task.ID)
}

func main() {
	fmt.Println("--- gprMax CI/CD Performance Runner (Go) ---")

	// Goroutine ka use karke simultaneous task chalana
	task1 := SimTask{ID: 101, Name: "Antenna_Model_A"}
	
	go runSimulation(task1) // 'go' keyword se concurrency aati hai

	fmt.Println("⏳ System is monitoring the simulation...")
	time.Sleep(3 * time.Second) // Wait for goroutine to finish
	fmt.Println("🎯 All tasks synchronized. Ready for PR!")
}