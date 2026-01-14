# Go Interview Preparation Programs

This repository contains 16 comprehensive Go programs designed to help you prepare for Go/Golang interviews. Each program demonstrates key concepts and features of the Go programming language.

## How to Run the Programs

### Prerequisites
- Go 1.16 or later installed on your system
- Basic understanding of command line/terminal

### Running Individual Programs
Navigate to the directory containing the Go files and run each program using the `go run` command:

```bash
# Basic programs
go run 01_hello_world.go
go run 02_data_types.go  
go run 03_zero_values.go
go run 04_arrays_slices.go
go run 05_type_conversions.go

# Advanced concepts  
go run 06_error_handling.go
go run 07_goroutines_basic.go
go run 08_channels_basic.go
go run 09_channels_select.go
go run 10_range_keyword.go

# Object-oriented concepts
go run 11_structs_methods.go
go run 12_interfaces.go
go run 13_maps.go
go run 14_pointers.go

# Advanced topics
go run 15_concurrency_patterns.go
go run 16_memory_management.go
```

### Building Executables
You can also build executable files:

```bash
# Build a specific program
go build 01_hello_world.go
./01_hello_world

# Build all programs at once
for file in *.go; do
    go build "$file"
done
```

## Program Descriptions

### 1. Hello World (`01_hello_world.go`)
- Basic Go program structure
- Package declaration
- Import statements
- Main function

**Key Interview Topics:**
- Go program structure
- Package system
- Entry point of Go programs

### 2. Data Types (`02_data_types.go`)
- All basic data types in Go
- Numeric types (int, float, complex)
- String and rune types
- Boolean type
- Type information using %T

**Key Interview Topics:**
- Go's type system
- Difference between types
- Type safety

### 3. Zero Values (`03_zero_values.go`)
- Zero values for all data types
- Composite types zero values
- Nil values for pointers, slices, maps

**Key Interview Topics:**
- Memory initialization
- Default values
- Nil concept in Go

### 4. Arrays vs Slices (`04_arrays_slices.go`)
- Fixed-size arrays vs dynamic slices
- Slice operations (append, copy)
- Capacity vs length
- Slicing operations

**Key Interview Topics:**
- Memory allocation differences
- When to use arrays vs slices
- Slice internals

### 5. Type Conversions (`05_type_conversions.go`)
- Numeric type conversions
- String conversions
- Type assertions
- Type switches
- Interface{} usage

**Key Interview Topics:**
- Type safety and conversions
- Interface{} and type assertions
- Runtime type checking

### 6. Error Handling (`06_error_handling.go`)
- Basic error handling patterns
- Custom error types
- Error wrapping
- Multiple return values

**Key Interview Topics:**
- Go's approach to error handling
- Error vs exceptions
- Best practices for error handling

### 7. Basic Goroutines (`07_goroutines_basic.go`)
- Creating goroutines
- Concurrent vs sequential execution
- Anonymous goroutines
- Goroutine lifecycle

**Key Interview Topics:**
- Concurrency in Go
- Goroutines vs threads
- Go scheduler

### 8. Basic Channels (`08_channels_basic.go`)
- Unbuffered and buffered channels
- Channel directions (send-only, receive-only)
- Closing channels
- Channel communication patterns

**Key Interview Topics:**
- Channel types and behavior
- Communication between goroutines
- Channel blocking behavior

### 9. Channel Select (`09_channels_select.go`)
- Select statement
- Non-blocking operations
- Timeouts with channels
- Multiple channel communication

**Key Interview Topics:**
- Advanced channel patterns
- Non-blocking I/O
- Timeout handling

### 10. Range Keyword (`10_range_keyword.go`)
- Range over slices, arrays, maps
- Range over strings (runes)
- Range over channels
- Modifying elements during iteration

**Key Interview Topics:**
- Iteration patterns in Go
- Range behavior with different types
- Common pitfalls with range

### 11. Structs and Methods (`11_structs_methods.go`)
- Struct definition and initialization
- Methods with value and pointer receivers
- Struct embedding
- Anonymous structs

**Key Interview Topics:**
- Object-oriented programming in Go
- Value vs pointer receivers
- Composition over inheritance

### 12. Interfaces (`12_interfaces.go`)
- Interface definition and implementation
- Interface composition
- Type assertion and type switches
- Empty interface

**Key Interview Topics:**
- Interface implementation (implicit)
- Polymorphism in Go
- Interface best practices

### 13. Maps (`13_maps.go`)
- Map creation and operations
- Iterating over maps
- Nested maps
- Map comparison and nil maps

**Key Interview Topics:**
- Hash table implementation
- Map characteristics
- Memory considerations with maps

### 14. Pointers (`14_pointers.go`)
- Pointer basics and operations
- Pointers with functions
- Pointers with structs and slices
- Memory addressing

**Key Interview Topics:**
- Memory management
- Pass by value vs reference
- Pointer arithmetic limitations

### 15. Concurrency Patterns (`15_concurrency_patterns.go`)
- Worker pool pattern
- Fan-in/Fan-out pattern
- Pipeline pattern
- WaitGroup and Mutex usage

**Key Interview Topics:**
- Advanced concurrency patterns
- Synchronization primitives
- Real-world concurrency scenarios

### 16. Memory Management (`16_memory_management.go`)
- Garbage collection
- Memory allocation patterns
- Stack vs heap allocation
- Memory leak prevention

**Key Interview Topics:**
- Go's garbage collector
- Memory optimization
- Performance considerations

## Common Interview Questions Covered

1. **Basic Go Concepts:**
   - What is Go and why was it created?
   - Go workspace and GOPATH
   - Package system and imports

2. **Data Types and Variables:**
   - Go's type system
   - Zero values
   - Type conversions and assertions

3. **Control Structures:**
   - Range keyword
   - Type switches
   - Select statements

4. **Functions and Methods:**
   - Function declarations
   - Method receivers
   - Error handling patterns

5. **Concurrency:**
   - Goroutines vs threads
   - Channels and communication
   - Synchronization primitives

6. **Memory Management:**
   - Garbage collection
   - Pointers and memory allocation
   - Performance optimization

7. **Advanced Topics:**
   - Interfaces and polymorphism
   - Composition patterns
   - Concurrency patterns

## Tips for Interview Preparation

1. **Run each program** and understand the output
2. **Modify the code** to see how behavior changes
3. **Practice explaining** the concepts out loud
4. **Focus on the differences** between Go and other languages
5. **Understand the trade-offs** of different approaches
6. **Practice coding** similar programs from scratch

## Additional Practice

Try modifying these programs to:
- Add error handling where missing
- Implement additional methods
- Create more complex examples
- Combine concepts from multiple programs
- Optimize for performance

Good luck with your Go interview preparation!
