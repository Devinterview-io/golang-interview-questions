// 11_structs_methods.go - Demonstrates structs and methods
package main

import "fmt"

// Basic struct definition
type Person struct {
	Name    string
	Age     int
	Email   string
	Address Address // Embedded struct
}

// Nested struct
type Address struct {
	Street  string
	City    string
	Country string
}

// Method with value receiver
func (p Person) GetFullInfo() string {
	return fmt.Sprintf("%s (%d years old) - %s", p.Name, p.Age, p.Email)
}

// Method with pointer receiver (can modify the struct)
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

// Method with pointer receiver
func (p *Person) UpdateEmail(newEmail string) {
	p.Email = newEmail
}

// Method on Address struct
func (a Address) GetFullAddress() string {
	return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.Country)
}

func main() {
	fmt.Println("=== STRUCTS AND METHODS DEMONSTRATION ===")

	// Creating structs
	fmt.Println("--- CREATING STRUCTS ---")

	// Method 1: Field by field
	var person1 Person
	person1.Name = "Alice"
	person1.Age = 30
	person1.Email = "alice@example.com"
	person1.Address = Address{
		Street:  "123 Main St",
		City:    "New York",
		Country: "USA",
	}

	// Method 2: Struct literal
	person2 := Person{
		Name:  "Bob",
		Age:   25,
		Email: "bob@example.com",
		Address: Address{
			Street:  "456 Oak Ave",
			City:    "Los Angeles",
			Country: "USA",
		},
	}

	// Method 3: Using field names (order doesn't matter)
	person3 := Person{
		Email: "charlie@example.com",
		Name:  "Charlie",
		Age:   35,
		Address: Address{
			Country: "Canada",
			City:    "Toronto",
			Street:  "789 Pine Rd",
		},
	}

	// Method 4: Positional (not recommended for readability)
	address4 := Address{"321 Elm St", "Chicago", "USA"}
	person4 := Person{"Diana", 28, "diana@example.com", address4}

	fmt.Printf("Person 1: %+v\n", person1)
	fmt.Printf("Person 2: %+v\n", person2)
	fmt.Printf("Person 3: %+v\n", person3)
	fmt.Printf("Person 4: %+v\n", person4)

	// Using methods
	fmt.Println("\n--- USING METHODS ---")
	fmt.Printf("Person 1 info: %s\n", person1.GetFullInfo())
	fmt.Printf("Person 1 address: %s\n", person1.Address.GetFullAddress())

	// Method with pointer receiver
	fmt.Println("\n--- POINTER RECEIVER METHODS ---")
	fmt.Printf("Before update - %s\n", person2.GetFullInfo())

	person2.UpdateAge(26)
	person2.UpdateEmail("bob.updated@example.com")

	fmt.Printf("After update - %s\n", person2.GetFullInfo())

	// Anonymous structs
	fmt.Println("\n--- ANONYMOUS STRUCTS ---")
	car := struct {
		Brand string
		Model string
		Year  int
	}{
		Brand: "Toyota",
		Model: "Camry",
		Year:  2023,
	}

	fmt.Printf("Car: %+v\n", car)

	// Struct comparison
	fmt.Println("\n--- STRUCT COMPARISON ---")
	addr1 := Address{"123 Main St", "New York", "USA"}
	addr2 := Address{"123 Main St", "New York", "USA"}
	addr3 := Address{"456 Oak Ave", "Los Angeles", "USA"}

	fmt.Printf("addr1 == addr2: %t\n", addr1 == addr2)
	fmt.Printf("addr1 == addr3: %t\n", addr1 == addr3)

	// Struct embedding (composition)
	fmt.Println("\n--- STRUCT EMBEDDING ---")
	type Employee struct {
		Person     // Embedded struct
		Position   string
		Salary     float64
		Department string
	}

	employee := Employee{
		Person: Person{
			Name:  "John",
			Age:   32,
			Email: "john@company.com",
			Address: Address{
				Street:  "100 Business Blvd",
				City:    "San Francisco",
				Country: "USA",
			},
		},
		Position:   "Software Engineer",
		Salary:     95000.0,
		Department: "Engineering",
	}

	fmt.Printf("Employee: %+v\n", employee)
	fmt.Printf("Employee name (direct access): %s\n", employee.Name) // Can access embedded fields directly
	fmt.Printf("Employee info: %s\n", employee.GetFullInfo())        // Can call embedded methods directly

	// Struct tags (used with JSON, XML, etc.)
	fmt.Println("\n--- STRUCT TAGS ---")
	type Product struct {
		ID    int     `json:"id" xml:"id"`
		Name  string  `json:"name" xml:"name"`
		Price float64 `json:"price" xml:"price"`
	}

	product := Product{ID: 1, Name: "Laptop", Price: 999.99}
	fmt.Printf("Product with tags: %+v\n", product)
}
