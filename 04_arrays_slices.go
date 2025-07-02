// 04_arrays_slices.go - Demonstrates arrays vs slices
package main

import "fmt"

func main() {
	fmt.Println("=== ARRAYS VS SLICES DEMONSTRATION ===")

	// Arrays - Fixed size
	fmt.Println("--- ARRAYS ---")
	var arr1 [5]int               // Declaration with zero values
	arr2 := [5]int{1, 2, 3, 4, 5} // Declaration with initialization
	arr3 := [...]int{10, 20, 30}  // Let compiler count the elements

	fmt.Printf("arr1: %v (length: %d)\n", arr1, len(arr1))
	fmt.Printf("arr2: %v (length: %d)\n", arr2, len(arr2))
	fmt.Printf("arr3: %v (length: %d)\n", arr3, len(arr3))

	// Modifying array elements
	arr1[0] = 100
	fmt.Printf("arr1 after modification: %v\n", arr1)

	// Slices - Dynamic size
	fmt.Println("\n--- SLICES ---")
	var slice1 []int               // nil slice
	slice2 := []int{1, 2, 3, 4, 5} // slice literal
	slice3 := make([]int, 5)       // make function with length
	slice4 := make([]int, 3, 10)   // make with length and capacity

	fmt.Printf("slice1: %v (len: %d, cap: %d, nil: %t)\n", slice1, len(slice1), cap(slice1), slice1 == nil)
	fmt.Printf("slice2: %v (len: %d, cap: %d)\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3: %v (len: %d, cap: %d)\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("slice4: %v (len: %d, cap: %d)\n", slice4, len(slice4), cap(slice4))

	// Slicing operations
	fmt.Println("\n--- SLICING OPERATIONS ---")
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("numbers[2:5]: %v\n", numbers[2:5])
	fmt.Printf("numbers[:4]: %v\n", numbers[:4])
	fmt.Printf("numbers[6:]: %v\n", numbers[6:])
	fmt.Printf("numbers[:]: %v\n", numbers[:])

	// Append operations
	fmt.Println("\n--- APPEND OPERATIONS ---")
	var dynamicSlice []int
	fmt.Printf("Initial: %v (len: %d, cap: %d)\n", dynamicSlice, len(dynamicSlice), cap(dynamicSlice))

	dynamicSlice = append(dynamicSlice, 1)
	fmt.Printf("After append(1): %v (len: %d, cap: %d)\n", dynamicSlice, len(dynamicSlice), cap(dynamicSlice))

	dynamicSlice = append(dynamicSlice, 2, 3, 4)
	fmt.Printf("After append(2,3,4): %v (len: %d, cap: %d)\n", dynamicSlice, len(dynamicSlice), cap(dynamicSlice))

	moreNumbers := []int{5, 6, 7}
	dynamicSlice = append(dynamicSlice, moreNumbers...)
	fmt.Printf("After append(slice...): %v (len: %d, cap: %d)\n", dynamicSlice, len(dynamicSlice), cap(dynamicSlice))

	// Copy operations
	fmt.Println("\n--- COPY OPERATIONS ---")
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	n := copy(dst, src)
	fmt.Printf("Source: %v\n", src)
	fmt.Printf("Destination: %v\n", dst)
	fmt.Printf("Elements copied: %d\n", n)
}
