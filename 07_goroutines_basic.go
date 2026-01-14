// 07_goroutines_basic.go - Basic goroutines demonstration
package main

import (
	"fmt"
	"runtime"
	"time"
)

func sayHello(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Hello from %s - %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func countNumbers(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s counting: %d\n", name, i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	fmt.Println("=== GOROUTINES BASIC DEMONSTRATION ===")

	// Show initial number of goroutines
	fmt.Printf("Initial goroutines: %d\n", runtime.NumGoroutine())

	// Sequential execution (normal function calls)
	fmt.Println("\n--- SEQUENTIAL EXECUTION ---")
	start := time.Now()
	sayHello("Sequential-1")
	sayHello("Sequential-2")
	fmt.Printf("Sequential execution took: %v\n", time.Since(start))

	// Concurrent execution using goroutines
	fmt.Println("\n--- CONCURRENT EXECUTION ---")
	start = time.Now()

	go sayHello("Goroutine-1")
	go sayHello("Goroutine-2")
	go countNumbers("Counter")

	// Show number of goroutines after launching
	fmt.Printf("Goroutines after launch: %d\n", runtime.NumGoroutine())

	// Wait for goroutines to complete
	time.Sleep(1 * time.Second)

	fmt.Printf("Concurrent execution took: %v\n", time.Since(start))
	fmt.Printf("Final goroutines: %d\n", runtime.NumGoroutine())

	// Anonymous goroutines
	fmt.Println("\n--- ANONYMOUS GOROUTINES ---")

	go func() {
		fmt.Println("Anonymous goroutine 1 started")
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Anonymous goroutine 1 finished")
	}()

	go func(msg string) {
		fmt.Printf("Anonymous goroutine 2: %s\n", msg)
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Anonymous goroutine 2 finished")
	}("Hello from closure!")

	// Wait for anonymous goroutines
	time.Sleep(500 * time.Millisecond)

	fmt.Println("Main function ending...")
}
