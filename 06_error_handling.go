// 06_error_handling.go - Demonstrates error handling patterns in Go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Custom error type
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error in field '%s': %s", e.Field, e.Message)
}

// Function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Function with custom error
func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Message: "age cannot be negative",
		}
	}
	if age > 150 {
		return &ValidationError{
			Field:   "age",
			Message: "age seems unrealistic",
		}
	}
	return nil
}

// Function that wraps errors
func parseAndValidateAge(ageStr string) (int, error) {
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return 0, fmt.Errorf("failed to parse age '%s': %w", ageStr, err)
	}

	if err := validateAge(age); err != nil {
		return 0, fmt.Errorf("age validation failed: %w", err)
	}

	return age, nil
}

func main() {
	fmt.Println("=== ERROR HANDLING DEMONSTRATION ===")

	// Basic error handling
	fmt.Println("--- BASIC ERROR HANDLING ---")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result)
	}

	// Custom error handling
	fmt.Println("\n--- CUSTOM ERROR HANDLING ---")
	testAges := []int{25, -5, 200}
	for _, age := range testAges {
		if err := validateAge(age); err != nil {
			fmt.Printf("Age %d: %v\n", age, err)

			// Type assertion to check for custom error
			if ve, ok := err.(*ValidationError); ok {
				fmt.Printf("  Custom error - Field: %s, Message: %s\n", ve.Field, ve.Message)
			}
		} else {
			fmt.Printf("Age %d: valid\n", age)
		}
	}

	// Error wrapping
	fmt.Println("\n--- ERROR WRAPPING ---")
	testAgeStrings := []string{"25", "abc", "-10", "200"}
	for _, ageStr := range testAgeStrings {
		age, err := parseAndValidateAge(ageStr)
		if err != nil {
			fmt.Printf("Age string '%s': %v\n", ageStr, err)

			// Unwrap the error to check the underlying cause
			if unwrapped := errors.Unwrap(err); unwrapped != nil {
				fmt.Printf("  Underlying error: %v\n", unwrapped)
			}
		} else {
			fmt.Printf("Age string '%s': parsed to %d\n", ageStr, age)
		}
	}

	// Multiple return values for error handling
	fmt.Println("\n--- MULTIPLE RETURN VALUES ---")
	value, success := safeDivide(20, 4)
	if success {
		fmt.Printf("20 / 4 = %.2f\n", value)
	} else {
		fmt.Println("Division failed")
	}

	value, success = safeDivide(20, 0)
	if success {
		fmt.Printf("20 / 0 = %.2f\n", value)
	} else {
		fmt.Println("Division failed")
	}
}

// Alternative error handling using boolean
func safeDivide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}
