// 14_pointers.go - Demonstrates pointers in Go
package main

import "fmt"

func main() {
	fmt.Println("=== POINTERS DEMONSTRATION ===")

	// Basic pointer operations
	fmt.Println("--- BASIC POINTER OPERATIONS ---")

	var x int = 42
	var p *int = &x // p points to x

	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("Value of p (address): %p\n", p)
	fmt.Printf("Value pointed to by p: %d\n", *p)

	// Modify value through pointer
	*p = 100
	fmt.Printf("After *p = 100, x = %d\n", x)

	// Zero value of pointer is nil
	fmt.Println("\n--- NIL POINTERS ---")
	var nilPtr *int
	fmt.Printf("Nil pointer: %v (nil: %t)\n", nilPtr, nilPtr == nil)

	// Creating pointers using new
	fmt.Println("\n--- USING NEW ---")
	ptr := new(int)
	*ptr = 50
	fmt.Printf("Pointer from new: %p, value: %d\n", ptr, *ptr)

	// Pointers with functions
	fmt.Println("\n--- POINTERS WITH FUNCTIONS ---")
	a := 10
	b := 20

	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)
	swapByValue(a, b)
	fmt.Printf("After swapByValue: a=%d, b=%d\n", a, b)

	swapByPointer(&a, &b)
	fmt.Printf("After swapByPointer: a=%d, b=%d\n", a, b)

	// Pointers with structs
	fmt.Println("\n--- POINTERS WITH STRUCTS ---")
	type Person struct {
		Name string
		Age  int
	}

	person1 := Person{"Alice", 25}
	fmt.Printf("person1: %+v\n", person1)

	// Pointer to struct
	personPtr := &person1
	fmt.Printf("Access via pointer: %+v\n", *personPtr)

	// Go automatically dereferences struct pointers
	fmt.Printf("Name via pointer: %s\n", personPtr.Name)
	personPtr.Age = 26
	fmt.Printf("After updating age via pointer: %+v\n", person1)

	// Creating struct pointer directly
	person2 := &Person{"Bob", 30}
	fmt.Printf("person2 (created as pointer): %+v\n", *person2)

	// Pointers with slices
	fmt.Println("\n--- POINTERS WITH SLICES ---")
	slice1 := []int{1, 2, 3}
	slice2 := slice1 // Slices are reference types

	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)

	slice2[0] = 10
	fmt.Printf("After modifying slice2[0]: slice1=%v, slice2=%v\n", slice1, slice2)

	// Explicit pointer to slice
	slicePtr := &slice1
	(*slicePtr)[1] = 20
	fmt.Printf("After modifying via slicePtr: %v\n", slice1)

	// Pointers with arrays
	fmt.Println("\n--- POINTERS WITH ARRAYS ---")
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1 // Arrays are value types (copied)

	fmt.Printf("arr1: %v\n", arr1)
	fmt.Printf("arr2: %v\n", arr2)

	arr2[0] = 10
	fmt.Printf("After modifying arr2[0]: arr1=%v, arr2=%v\n", arr1, arr2)

	// Pointer to array
	arrPtr := &arr1
	(*arrPtr)[1] = 20
	fmt.Printf("After modifying via arrPtr: %v\n", arr1)

	// Pointer arithmetic is not allowed in Go
	fmt.Println("\n--- POINTER LIMITATIONS ---")
	nums := []int{1, 2, 3, 4, 5}
	ptr1 := &nums[0]
	ptr2 := &nums[1]

	fmt.Printf("ptr1 points to: %d\n", *ptr1)
	fmt.Printf("ptr2 points to: %d\n", *ptr2)

	// This would not compile:
	// ptr1++           // Pointer arithmetic not allowed
	// ptr3 := ptr1 + 1 // Pointer arithmetic not allowed

	// Double pointers
	fmt.Println("\n--- DOUBLE POINTERS ---")
	value := 42
	ptr3 := &value
	doublePtr := &ptr3

	fmt.Printf("value: %d\n", value)
	fmt.Printf("*ptr3: %d\n", *ptr3)
	fmt.Printf("**doublePtr: %d\n", **doublePtr)

	**doublePtr = 100
	fmt.Printf("After **doublePtr = 100, value = %d\n", value)
}

// Function that receives values (won't modify originals)
func swapByValue(x, y int) {
	x, y = y, x
	fmt.Printf("Inside swapByValue: x=%d, y=%d\n", x, y)
}

// Function that receives pointers (can modify originals)
func swapByPointer(x, y *int) {
	*x, *y = *y, *x
	fmt.Printf("Inside swapByPointer: *x=%d, *y=%d\n", *x, *y)
}
