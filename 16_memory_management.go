// 16_memory_management.go - Demonstrates memory management concepts
package main

import (
	"fmt"
	"runtime"
	"time"
)

type LargeStruct struct {
	Data [1000000]int // Large array to consume memory
	ID   int
}

func main() {
	fmt.Println("=== MEMORY MANAGEMENT DEMONSTRATION ===")

	// Show initial memory stats
	showMemStats("Initial")

	// Create many objects to trigger garbage collection
	fmt.Println("\n--- CREATING LARGE OBJECTS ---")
	createLargeObjects()

	showMemStats("After creating objects")

	// Force garbage collection
	fmt.Println("\n--- FORCING GARBAGE COLLECTION ---")
	runtime.GC()
	showMemStats("After manual GC")

	// Memory allocation patterns
	fmt.Println("\n--- MEMORY ALLOCATION PATTERNS ---")
	demonstrateAllocationPatterns()

	// Stack vs Heap allocation
	fmt.Println("\n--- STACK VS HEAP ALLOCATION ---")
	demonstrateStackVsHeap()

	// Memory leaks prevention
	fmt.Println("\n--- MEMORY LEAKS PREVENTION ---")
	demonstrateMemoryLeakPrevention()
}

func showMemStats(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("%s Memory Stats:\n", label)
	fmt.Printf("  Alloc = %d KB", bToKb(m.Alloc))
	fmt.Printf("  TotalAlloc = %d KB", bToKb(m.TotalAlloc))
	fmt.Printf("  Sys = %d KB", bToKb(m.Sys))
	fmt.Printf("  NumGC = %d\n", m.NumGC)
}

func bToKb(b uint64) uint64 {
	return b / 1024
}

func createLargeObjects() {
	objects := make([]*LargeStruct, 100)

	for i := 0; i < 100; i++ {
		objects[i] = &LargeStruct{ID: i}
		if i%20 == 0 {
			fmt.Printf("Created object %d\n", i)
			showMemStats(fmt.Sprintf("Object %d", i))
		}
	}

	// Objects will be eligible for GC when function returns
	fmt.Printf("Created %d large objects\n", len(objects))
}

func demonstrateAllocationPatterns() {
	// Slice allocation and growth
	var slice []int

	for i := 0; i < 1000; i++ {
		slice = append(slice, i)
		if i == 0 || i == 1 || i == 2 || i == 4 || i == 8 || i == 16 || i == 32 || i == 64 || i == 128 || i == 256 || i == 512 {
			fmt.Printf("Slice length: %d, capacity: %d\n", len(slice), cap(slice))
		}
	}

	// Pre-allocating vs growing
	fmt.Println("\nPre-allocated slice:")
	preAllocated := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		preAllocated = append(preAllocated, i)
	}
	fmt.Printf("Pre-allocated - length: %d, capacity: %d\n", len(preAllocated), cap(preAllocated))
}

func demonstrateStackVsHeap() {
	// Stack allocation (typically)
	stackVar := 42
	fmt.Printf("Stack variable address: %p\n", &stackVar)

	// Heap allocation (escapes to heap)
	heapVar := createHeapVariable()
	fmt.Printf("Heap variable address: %p\n", heapVar)

	// Large local variables might escape to heap
	largeArray := [100000]int{}
	fmt.Printf("Large array address: %p\n", &largeArray)
}

func createHeapVariable() *int {
	x := 42
	return &x // x escapes to heap because we return its address
}

func demonstrateMemoryLeakPrevention() {
	// Channel without proper cleanup (potential leak)
	ch := make(chan int, 100)

	// Good practice: always close channels when done
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // Important: close the channel
	}()

	// Consume all values
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}

	// Slice memory management
	demonstrateSliceMemoryManagement()

	// Goroutine cleanup
	demonstrateGoroutineCleanup()
}

func demonstrateSliceMemoryManagement() {
	// Creating a large slice
	largeSlice := make([]int, 1000000)
	for i := range largeSlice {
		largeSlice[i] = i
	}

	// Taking a small sub-slice - this keeps the entire backing array in memory
	badSubSlice := largeSlice[0:10]

	// Good practice: copy the data if you only need a small portion
	goodSubSlice := make([]int, 10)
	copy(goodSubSlice, largeSlice[0:10])

	fmt.Printf("Bad sub-slice length: %d, capacity: %d\n", len(badSubSlice), cap(badSubSlice))
	fmt.Printf("Good sub-slice length: %d, capacity: %d\n", len(goodSubSlice), cap(goodSubSlice))

	// Now largeSlice can be garbage collected if goodSubSlice is used instead of badSubSlice
	_ = goodSubSlice
	largeSlice = nil // Help GC by explicitly setting to nil
}

func demonstrateGoroutineCleanup() {
	done := make(chan bool)

	// Start a goroutine with proper cleanup
	go func() {
		defer func() {
			fmt.Println("Goroutine cleanup completed")
			done <- true
		}()

		// Simulate some work
		for i := 0; i < 5; i++ {
			fmt.Printf("Goroutine working: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Wait for goroutine to complete
	<-done
	fmt.Println("Goroutine finished properly")
}
