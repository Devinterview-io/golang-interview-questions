# Quick Reference Guide - Go Interview Programs

## Run Commands Summary

```bash
# Individual programs
go run 01_hello_world.go           # Basic Go structure
go run 02_data_types.go            # All data types
go run 03_zero_values.go           # Zero values concept
go run 04_arrays_slices.go         # Arrays vs Slices
go run 05_type_conversions.go      # Type conversions & assertions
go run 06_error_handling.go        # Error handling patterns
go run 07_goroutines_basic.go      # Basic goroutines
go run 08_channels_basic.go        # Channel fundamentals
go run 09_channels_select.go       # Select statement
go run 10_range_keyword.go         # Range iteration
go run 11_structs_methods.go       # Structs and methods
go run 12_interfaces.go            # Interface implementation
go run 13_maps.go                  # Map operations
go run 14_pointers.go              # Pointer concepts
go run 15_concurrency_patterns.go  # Advanced concurrency
go run 16_memory_management.go     # Memory management

# Run all programs interactively
./run_all_programs.sh

# Test compilation of all programs
./test_compilation.sh
```

## Key Concepts Covered

### Fundamentals
- ✅ Package system and imports
- ✅ Data types and zero values
- ✅ Variables and constants
- ✅ Type system and conversions

### Data Structures
- ✅ Arrays and slices
- ✅ Maps and their operations
- ✅ Structs and methods
- ✅ Pointers and memory addressing

### Object-Oriented Programming
- ✅ Structs and methods
- ✅ Interfaces and implementation
- ✅ Composition over inheritance
- ✅ Polymorphism through interfaces

### Concurrency
- ✅ Goroutines basics
- ✅ Channels and communication
- ✅ Select statements
- ✅ Advanced concurrency patterns
- ✅ Synchronization (WaitGroup, Mutex)

### Advanced Topics
- ✅ Error handling patterns
- ✅ Memory management
- ✅ Garbage collection
- ✅ Range keyword usage

### Best Practices
- ✅ Proper error handling
- ✅ Memory leak prevention
- ✅ Concurrency patterns
- ✅ Interface design

## Interview Tips

1. **Run each program** and understand the output completely
2. **Practice explaining** concepts while the program runs
3. **Modify programs** to test your understanding
4. **Focus on Go-specific** features like goroutines, channels, interfaces
5. **Understand the philosophy** behind Go's design decisions

## Common Interview Questions Mapped to Programs

| Question | Program |
|----------|---------|
| "Explain Go's type system" | 02, 03, 05 |
| "Difference between arrays and slices" | 04 |
| "How does error handling work in Go?" | 06 |
| "Explain goroutines and channels" | 07, 08, 09 |
| "How do interfaces work in Go?" | 12 |
| "Explain Go's memory management" | 16 |
| "What are some concurrency patterns?" | 15 |
| "How do pointers work in Go?" | 14 |

## Quick Test
Run this to verify everything works:
```bash
./test_compilation.sh && echo "✅ All programs ready for interview prep!"
```
