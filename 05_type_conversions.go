// 05_type_conversions.go - Demonstrates type conversions and assertions
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== TYPE CONVERSIONS DEMONSTRATION ===")

	// Numeric conversions
	fmt.Println("--- NUMERIC CONVERSIONS ---")
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("int: %d -> float64: %f -> uint: %d\n", i, f, u)

	// String conversions
	fmt.Println("\n--- STRING CONVERSIONS ---")

	// Number to string
	num := 123
	str1 := strconv.Itoa(num)
	str2 := fmt.Sprintf("%d", num)

	fmt.Printf("Number %d to string: '%s' (using Itoa)\n", num, str1)
	fmt.Printf("Number %d to string: '%s' (using Sprintf)\n", num, str2)

	// String to number
	str := "456"
	num2, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Error converting string to int: %v\n", err)
	} else {
		fmt.Printf("String '%s' to number: %d\n", str, num2)
	}

	// Float conversions
	floatStr := "3.14159"
	floatNum, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Printf("Error converting string to float: %v\n", err)
	} else {
		fmt.Printf("String '%s' to float: %f\n", floatStr, floatNum)
	}

	// Boolean conversions
	boolStr := "true"
	boolVal, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Printf("Error converting string to bool: %v\n", err)
	} else {
		fmt.Printf("String '%s' to bool: %t\n", boolStr, boolVal)
	}

	// Type assertions with interfaces
	fmt.Println("\n--- TYPE ASSERTIONS ---")
	var data interface{} = "Hello, Go!"

	// Type assertion with ok pattern
	if str, ok := data.(string); ok {
		fmt.Printf("data is a string: '%s'\n", str)
	} else {
		fmt.Println("data is not a string")
	}

	// Type assertion without ok pattern (can panic)
	str3 := data.(string)
	fmt.Printf("Direct assertion: '%s'\n", str3)

	// Type switch
	fmt.Println("\n--- TYPE SWITCH ---")
	testTypeSwitch(42)
	testTypeSwitch("Hello")
	testTypeSwitch(3.14)
	testTypeSwitch(true)
	testTypeSwitch([]int{1, 2, 3})
}

func testTypeSwitch(data interface{}) {
	switch v := data.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: '%s'\n", v)
	case float64:
		fmt.Printf("Float64: %f\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T with value: %v\n", v, v)
	}
}
