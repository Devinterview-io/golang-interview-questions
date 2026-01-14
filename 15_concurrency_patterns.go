// 15_concurrency_patterns.go - Advanced concurrency patterns
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== CONCURRENCY PATTERNS DEMONSTRATION ===")

	// Worker Pool Pattern
	fmt.Println("--- WORKER POOL PATTERN ---")
	workerPoolDemo()

	// Fan-in Fan-out Pattern
	fmt.Println("\n--- FAN-IN FAN-OUT PATTERN ---")
	fanInOutDemo()

	// Pipeline Pattern
	fmt.Println("\n--- PIPELINE PATTERN ---")
	pipelineDemo()

	// WaitGroup Example
	fmt.Println("\n--- WAITGROUP EXAMPLE ---")
	waitGroupDemo()

	// Mutex Example
	fmt.Println("\n--- MUTEX EXAMPLE ---")
	mutexDemo()
}

// Worker Pool Pattern
func workerPoolDemo() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond) // Simulate work
		results <- job * 2
	}
}

// Fan-in Fan-out Pattern
func fanInOutDemo() {
	input := make(chan int)

	// Fan-out: distribute work to multiple goroutines
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go fanOut(input, ch1, ch2, ch3)

	// Process in parallel
	out1 := process("Process-1", ch1)
	out2 := process("Process-2", ch2)
	out3 := process("Process-3", ch3)

	// Fan-in: combine results
	combined := fanIn(out1, out2, out3)

	// Send some data
	go func() {
		for i := 1; i <= 6; i++ {
			input <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(input)
	}()

	// Collect results
	for result := range combined {
		fmt.Printf("Combined result: %d\n", result)
	}
}

func fanOut(input <-chan int, outputs ...chan<- int) {
	go func() {
		defer func() {
			for _, out := range outputs {
				close(out)
			}
		}()

		for value := range input {
			for i, out := range outputs {
				select {
				case out <- value:
					fmt.Printf("Sent %d to output %d\n", value, i+1)
				default:
					// Non-blocking send
				}
			}
		}
	}()
}

func process(name string, input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for value := range input {
			processed := value * value // Square the value
			fmt.Printf("%s processed %d -> %d\n", name, value, processed)
			output <- processed
			time.Sleep(50 * time.Millisecond)
		}
	}()
	return output
}

func fanIn(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(inputs))

	for _, input := range inputs {
		go func(ch <-chan int) {
			defer wg.Done()
			for value := range ch {
				output <- value
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

// Pipeline Pattern
func pipelineDemo() {
	// Stage 1: Generate numbers
	numbers := generateNumbers(5)

	// Stage 2: Square numbers
	squared := squareNumbers(numbers)

	// Stage 3: Sum numbers
	sum := sumNumbers(squared)

	fmt.Printf("Pipeline result: %d\n", <-sum)
}

func generateNumbers(count int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= count; i++ {
			out <- i
			fmt.Printf("Generated: %d\n", i)
		}
	}()
	return out
}

func squareNumbers(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range input {
			squared := num * num
			out <- squared
			fmt.Printf("Squared: %d -> %d\n", num, squared)
		}
	}()
	return out
}

func sumNumbers(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		sum := 0
		for num := range input {
			sum += num
			fmt.Printf("Sum so far: %d\n", sum)
		}
		out <- sum
	}()
	return out
}

// WaitGroup Example
func waitGroupDemo() {
	var wg sync.WaitGroup

	tasks := []string{"Task-1", "Task-2", "Task-3", "Task-4"}

	for _, task := range tasks {
		wg.Add(1)
		go func(taskName string) {
			defer wg.Done()
			fmt.Printf("Starting %s\n", taskName)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("Completed %s\n", taskName)
		}(task)
	}

	fmt.Println("Waiting for all tasks to complete...")
	wg.Wait()
	fmt.Println("All tasks completed!")
}

// Mutex Example
func mutexDemo() {
	var mutex sync.Mutex
	var counter int
	var wg sync.WaitGroup

	numGoroutines := 10
	incrementsPerGoroutine := 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}

	wg.Wait()
	expected := numGoroutines * incrementsPerGoroutine
	fmt.Printf("Expected: %d, Actual: %d, Match: %t\n", expected, counter, expected == counter)
}
