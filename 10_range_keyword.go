// 10_range_keyword.go - Demonstrates the range keyword
package main

import "fmt"

func main() {
	fmt.Println("=== RANGE KEYWORD DEMONSTRATION ===")

	// Range over arrays and slices
	fmt.Println("--- RANGE OVER SLICE ---")
	numbers := []int{10, 20, 30, 40, 50}

	// Range with both index and value
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Range with only value (ignore index with _)
	fmt.Println("\nUsing only values:")
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}

	// Range with only index
	fmt.Println("\nUsing only indices:")
	for index := range numbers {
		fmt.Printf("Index: %d\n", index)
	}

	// Range over maps
	fmt.Println("\n--- RANGE OVER MAP ---")
	person := map[string]interface{}{
		"name":    "Alice",
		"age":     30,
		"city":    "New York",
		"married": true,
	}

	for key, value := range person {
		fmt.Printf("Key: %s, Value: %v (Type: %T)\n", key, value, value)
	}

	// Range over strings (iterates over runes)
	fmt.Println("\n--- RANGE OVER STRING ---")
	text := "Hello, 世界"

	fmt.Println("Character by character:")
	for index, runeValue := range text {
		fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", index, runeValue, runeValue)
	}

	// Range over channels
	fmt.Println("\n--- RANGE OVER CHANNEL ---")
	ch := make(chan int)

	// Start a goroutine to send data
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i * i // Send squares
		}
		close(ch) // Important: close the channel when done
	}()

	// Range over channel (receives until channel is closed)
	fmt.Println("Receiving squares from channel:")
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}

	// Range with arrays
	fmt.Println("\n--- RANGE OVER ARRAY ---")
	colors := [4]string{"red", "green", "blue", "yellow"}

	for i, color := range colors {
		fmt.Printf("Color %d: %s\n", i, color)
	}

	// Range with slice of structs
	fmt.Println("\n--- RANGE OVER SLICE OF STRUCTS ---")
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	for index, person := range people {
		fmt.Printf("Person %d: %s (age %d)\n", index, person.Name, person.Age)
	}

	// Modifying slice elements (be careful with range)
	fmt.Println("\n--- MODIFYING ELEMENTS ---")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original: %v\n", nums)

	// This won't modify the original slice because range creates copies
	for _, num := range nums {
		num = num * 2
	}
	fmt.Printf("After range modification (won't work): %v\n", nums)

	// Correct way to modify elements
	for i := range nums {
		nums[i] = nums[i] * 2
	}
	fmt.Printf("After index-based modification: %v\n", nums)
}
