#!/bin/bash

# test_compilation.sh - Script to test if all Go programs compile successfully

echo "Testing compilation of all Go programs..."
echo "========================================="

programs=(
    "01_hello_world.go"
    "02_data_types.go"
    "03_zero_values.go"
    "04_arrays_slices.go"
    "05_type_conversions.go"
    "06_error_handling.go"
    "07_goroutines_basic.go"
    "08_channels_basic.go"
    "09_channels_select.go"
    "10_range_keyword.go"
    "11_structs_methods.go"
    "12_interfaces.go"
    "13_maps.go"
    "14_pointers.go"
    "15_concurrency_patterns.go"
    "16_memory_management.go"
)

success_count=0
total_count=${#programs[@]}

for program in "${programs[@]}"; do
    if [ -f "$program" ]; then
        echo -n "Testing $program... "
        if go build "$program" 2>/dev/null; then
            echo "✓ OK"
            rm -f "${program%.go}"  # Remove the compiled binary
            ((success_count++))
        else
            echo "✗ FAILED"
            echo "Compilation errors:"
            go build "$program"
        fi
    else
        echo "✗ $program not found!"
    fi
done

echo ""
echo "========================================="
echo "Compilation Results:"
echo "Success: $success_count/$total_count programs"

if [ $success_count -eq $total_count ]; then
    echo "🎉 All programs compile successfully!"
    exit 0
else
    echo "❌ Some programs have compilation errors."
    exit 1
fi
