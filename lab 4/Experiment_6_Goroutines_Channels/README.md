# Experiment 6: Concurrent Tasks Using Goroutines and Channels

## Objective
To understand and implement concurrency in Go using goroutines and typed channels, and synchronize concurrent execution using `sync.WaitGroup`.

## Concepts Used

1. **Goroutines (`go` keyword)**: Lightweight threads of execution managed by the Go runtime rather than the operating system.
2. **Channels (`chan`)**: Typed conduits through which concurrent goroutines communicate and synchronize without explicit locking or race conditions.
   - Channel directionality (`chan<-`) ensures functions only send data, enforcing compile-time safety.
   - Buffered channels (`make(chan Result, capacity)`) allow sending values without blocking until capacity is reached.
3. **Synchronization (`sync.WaitGroup`)**:
   - `Add(delta)` increments the counter for active goroutines.
   - `Done()` decrements the counter upon completion (using `defer wg.Done()`).
   - `Wait()` blocks until the counter reaches zero.
4. **Channel Closure & Iteration (`close()`, `range`)**: Closing channels after goroutines finish allows safe draining of all collected values using `for val := range ch`.

## Program Description

The program launches three independent goroutines concurrently:
- **Goroutine 1 (`calculateSquares`)**: Computes the square of a series of numbers (`1, 2, 3, 4, 5`) and sends each result to `squareChan`.
- **Goroutine 2 (`calculateCubes`)**: Computes the cube of a series of numbers (`1, 2, 3, 4, 5`) and sends each result to `cubeChan`.
- **Goroutine 3 (`generateFibonacci`)**: Computes the first 5 terms of the Fibonacci sequence and sends each result to `fibChan`.

Each goroutine logs its startup, individual calculation steps, and completion. All results are received and neatly formatted in `main()`.

## Commands to Run

To run the program:

```bash
cd Experiment_6_Goroutines_Channels
go run main.go
```

To format code:

```bash
gofmt -w main.go
```

To build an executable:

```bash
go build -o experiment6.exe main.go
./experiment6.exe
```

## Expected Output

```text
================================================================
  Lab 4 - Experiment 6: Concurrent Tasks Using Goroutines & Channels
================================================================

--- Launching 3 Goroutines Concurrently ---
[Goroutine 3 - Fibonacci] >>> Started
[Goroutine 1 - Squares]   >>> Started
[Goroutine 2 - Cubes]     >>> Started
[Goroutine 3 - Fibonacci] Generating term 1: Fibonacci = 0
[Goroutine 1 - Squares]   Calculating square: 1^2 = 1
[Goroutine 2 - Cubes]     Calculating cube: 1^3 = 1
[Goroutine 3 - Fibonacci] Generating term 2: Fibonacci = 1
[Goroutine 1 - Squares]   Calculating square: 2^2 = 4
[Goroutine 2 - Cubes]     Calculating cube: 2^3 = 8
[Goroutine 3 - Fibonacci] Generating term 3: Fibonacci = 1
[Goroutine 1 - Squares]   Calculating square: 3^2 = 9
[Goroutine 3 - Fibonacci] Generating term 4: Fibonacci = 2
[Goroutine 2 - Cubes]     Calculating cube: 3^3 = 27
[Goroutine 1 - Squares]   Calculating square: 4^2 = 16
[Goroutine 3 - Fibonacci] Generating term 5: Fibonacci = 3
[Goroutine 3 - Fibonacci] <<< Completed
[Goroutine 2 - Cubes]     Calculating cube: 4^3 = 64
[Goroutine 1 - Squares]   Calculating square: 5^2 = 25
[Goroutine 1 - Squares]   <<< Completed
[Goroutine 2 - Cubes]     Calculating cube: 5^3 = 125
[Goroutine 2 - Cubes]     <<< Completed

All goroutines have finished execution.

================================================================
  Results Collected From Channels in main()
================================================================

1. Squares of Numbers:
   Number: 1 -> Square: 1
   Number: 2 -> Square: 4
   Number: 3 -> Square: 9
   Number: 4 -> Square: 16
   Number: 5 -> Square: 25

2. Cubes of Numbers:
   Number: 1 -> Cube: 1
   Number: 2 -> Cube: 8
   Number: 3 -> Cube: 27
   Number: 4 -> Cube: 64
   Number: 5 -> Cube: 125

3. Fibonacci Sequence:
   Term 1   -> Fibonacci: 0
   Term 2   -> Fibonacci: 1
   Term 3   -> Fibonacci: 1
   Term 4   -> Fibonacci: 2
   Term 5   -> Fibonacci: 3

================================================================
  Experiment 6 Completed Successfully
================================================================
```
