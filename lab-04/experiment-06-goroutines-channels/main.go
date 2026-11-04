package main

import (
	"fmt"
	"sync"
	"time"
)

// Result represents a calculated value with its corresponding input or term index.
type Result struct {
	Input int
	Value int
}

// calculateSquares calculates the square of each number in the input slice.
// It sends the results to the squareChan channel.
func calculateSquares(numbers []int, squareChan chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[Goroutine 1 - Squares]   >>> Started")

	for _, n := range numbers {
		// Simulate work delay to demonstrate concurrent interleaving
		time.Sleep(100 * time.Millisecond)
		sq := n * n
		fmt.Printf("[Goroutine 1 - Squares]   Calculating square: %d^2 = %d\n", n, sq)
		squareChan <- Result{Input: n, Value: sq}
	}

	fmt.Println("[Goroutine 1 - Squares]   <<< Completed")
}

// calculateCubes calculates the cube of each number in the input slice.
// It sends the results to the cubeChan channel.
func calculateCubes(numbers []int, cubeChan chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[Goroutine 2 - Cubes]     >>> Started")

	for _, n := range numbers {
		// Slightly different sleep duration to highlight concurrent execution
		time.Sleep(120 * time.Millisecond)
		cube := n * n * n
		fmt.Printf("[Goroutine 2 - Cubes]     Calculating cube: %d^3 = %d\n", n, cube)
		cubeChan <- Result{Input: n, Value: cube}
	}

	fmt.Println("[Goroutine 2 - Cubes]     <<< Completed")
}

// generateFibonacci generates the first 'count' terms of the Fibonacci sequence.
// It sends the results to the fibChan channel.
func generateFibonacci(count int, fibChan chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[Goroutine 3 - Fibonacci] >>> Started")

	a, b := 0, 1
	for i := 0; i < count; i++ {
		time.Sleep(90 * time.Millisecond)
		fmt.Printf("[Goroutine 3 - Fibonacci] Generating term %d: Fibonacci = %d\n", i+1, a)
		fibChan <- Result{Input: i + 1, Value: a}
		a, b = b, a+b
	}

	fmt.Println("[Goroutine 3 - Fibonacci] <<< Completed")
}

func main() {
	fmt.Println("================================================================")
	fmt.Println("  Lab 4 - Experiment 6: Concurrent Tasks Using Goroutines & Channels")
	fmt.Println("================================================================")

	numbers := []int{1, 2, 3, 4, 5}
	fibCount := 5

	// Create appropriately typed buffered channels to collect results without blocking
	squareChan := make(chan Result, len(numbers))
	cubeChan := make(chan Result, len(numbers))
	fibChan := make(chan Result, fibCount)

	// sync.WaitGroup is used to wait for all three goroutines to finish execution
	var wg sync.WaitGroup

	fmt.Println("\n--- Launching 3 Goroutines Concurrently ---")
	wg.Add(3)

	// Goroutine 1: Calculate Squares
	go calculateSquares(numbers, squareChan, &wg)

	// Goroutine 2: Calculate Cubes
	go calculateCubes(numbers, cubeChan, &wg)

	// Goroutine 3: Generate Fibonacci numbers
	go generateFibonacci(fibCount, fibChan, &wg)

	// Block until all 3 goroutines have signaled wg.Done()
	wg.Wait()

	// Safely close the channels as all goroutines have finished sending
	close(squareChan)
	close(cubeChan)
	close(fibChan)

	fmt.Println("\nAll goroutines have finished execution.")

	// Collect and display all results from channels in main()
	fmt.Println("\n================================================================")
	fmt.Println("  Results Collected From Channels in main()")
	fmt.Println("================================================================")

	fmt.Println("\n1. Squares of Numbers:")
	for res := range squareChan {
		fmt.Printf("   Number: %d -> Square: %d\n", res.Input, res.Value)
	}

	fmt.Println("\n2. Cubes of Numbers:")
	for res := range cubeChan {
		fmt.Printf("   Number: %d -> Cube: %d\n", res.Input, res.Value)
	}

	fmt.Println("\n3. Fibonacci Sequence:")
	for res := range fibChan {
		fmt.Printf("   Term %d   -> Fibonacci: %d\n", res.Input, res.Value)
	}

	fmt.Println("\n================================================================")
	fmt.Println("  Experiment 6 Completed Successfully")
	fmt.Println("================================================================")
}
