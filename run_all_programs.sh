#!/bin/bash

# run_all_programs.sh - Script to run all Go programs sequentially

echo "========================================="
echo "Go Interview Preparation Programs"
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

for program in "${programs[@]}"; do
    if [ -f "$program" ]; then
        echo ""
        echo "========================================="
        echo "Running: $program"
        echo "========================================="
        go run "$program"
        echo ""
        echo "Press Enter to continue to next program..."
        read -r
    else
        echo "Warning: $program not found!"
    fi
done

echo ""
echo "========================================="
echo "All programs completed!"
echo "========================================="
