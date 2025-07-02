// 13_maps.go - Demonstrates maps in Go
package main

import "fmt"

func main() {
	fmt.Println("=== MAPS DEMONSTRATION ===")

	// Creating maps
	fmt.Println("--- CREATING MAPS ---")

	// Method 1: Using make
	grades := make(map[string]int)
	grades["Alice"] = 85
	grades["Bob"] = 92
	grades["Charlie"] = 78

	// Method 2: Map literal
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	// Method 3: Empty map literal
	inventory := map[string]int{}
	inventory["apples"] = 50
	inventory["bananas"] = 30

	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Colors: %v\n", colors)
	fmt.Printf("Inventory: %v\n", inventory)

	// Accessing map values
	fmt.Println("\n--- ACCESSING VALUES ---")
	aliceGrade := grades["Alice"]
	fmt.Printf("Alice's grade: %d\n", aliceGrade)

	// Safe access with ok pattern
	if grade, ok := grades["Diana"]; ok {
		fmt.Printf("Diana's grade: %d\n", grade)
	} else {
		fmt.Println("Diana not found in grades")
	}

	// Check if key exists
	_, exists := colors["yellow"]
	fmt.Printf("Yellow exists in colors: %t\n", exists)

	// Map operations
	fmt.Println("\n--- MAP OPERATIONS ---")

	// Adding/updating
	grades["Diana"] = 88
	grades["Alice"] = 90 // Update existing
	fmt.Printf("Updated grades: %v\n", grades)

	// Deleting
	delete(grades, "Bob")
	fmt.Printf("After deleting Bob: %v\n", grades)

	// Length
	fmt.Printf("Number of students: %d\n", len(grades))

	// Iterating over maps
	fmt.Println("\n--- ITERATING OVER MAPS ---")

	// Iterate over key-value pairs
	fmt.Println("Color codes:")
	for color, code := range colors {
		fmt.Printf("  %s: %s\n", color, code)
	}

	// Iterate over keys only
	fmt.Println("Students:")
	for student := range grades {
		fmt.Printf("  %s\n", student)
	}

	// Iterate over values only
	fmt.Println("Grade values:")
	for _, grade := range grades {
		fmt.Printf("  %d\n", grade)
	}

	// Nested maps
	fmt.Println("\n--- NESTED MAPS ---")
	students := map[string]map[string]interface{}{
		"Alice": {
			"age":    20,
			"major":  "Computer Science",
			"gpa":    3.8,
			"active": true,
		},
		"Bob": {
			"age":    22,
			"major":  "Mathematics",
			"gpa":    3.6,
			"active": false,
		},
	}

	fmt.Printf("Students data: %v\n", students)

	// Accessing nested map values
	if alice, ok := students["Alice"]; ok {
		if age, ok := alice["age"]; ok {
			fmt.Printf("Alice's age: %v\n", age)
		}
		if major, ok := alice["major"]; ok {
			fmt.Printf("Alice's major: %v\n", major)
		}
	}

	// Map with struct values
	fmt.Println("\n--- MAP WITH STRUCT VALUES ---")
	type Person struct {
		Name string
		Age  int
		City string
	}

	people := map[int]Person{
		1: {"Alice", 25, "New York"},
		2: {"Bob", 30, "Los Angeles"},
		3: {"Charlie", 35, "Chicago"},
	}

	fmt.Printf("People: %v\n", people)

	// Modifying struct in map
	person := people[1]
	person.Age = 26
	people[1] = person // Need to reassign the whole struct
	fmt.Printf("Updated Alice: %v\n", people[1])

	// Map comparison (maps are not comparable)
	fmt.Println("\n--- MAP COMPARISON ---")
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"a": 1, "b": 2}

	// This would cause compilation error:
	// fmt.Printf("Maps equal: %t\n", map1 == map2)

	// Manual comparison
	equal := compareMaps(map1, map2)
	fmt.Printf("Maps are equal: %t\n", equal)

	// Zero value of map is nil
	fmt.Println("\n--- NIL MAPS ---")
	var nilMap map[string]int
	fmt.Printf("Nil map: %v (nil: %t)\n", nilMap, nilMap == nil)
	fmt.Printf("Length of nil map: %d\n", len(nilMap))

	// Reading from nil map returns zero value
	value := nilMap["key"]
	fmt.Printf("Value from nil map: %d\n", value)

	// Writing to nil map would panic
	// nilMap["key"] = 1 // This would panic
}

func compareMaps(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value1 := range map1 {
		if value2, ok := map2[key]; !ok || value1 != value2 {
			return false
		}
	}

	return true
}
