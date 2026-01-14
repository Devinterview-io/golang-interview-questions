// 02_data_types.go - Demonstrates all basic data types in Go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== GO DATA TYPES DEMONSTRATION ===")

	// Numeric Types
	var myInt int = 42
	var myInt8 int8 = 127
	var myInt16 int16 = 32767
	var myInt32 int32 = 2147483647
	var myInt64 int64 = 9223372036854775807

	var myUint uint = 42
	var myUint8 uint8 = 255
	var myByte byte = 255 // byte is alias for uint8

	var myFloat32 float32 = 3.14
	var myFloat64 float64 = 3.141592653589793

	var myComplex64 complex64 = 1 + 2i
	var myComplex128 complex128 = 1 + 2i

	// String and Rune
	var myString string = "Hello, Go!"
	var myRune rune = 'A' // rune is alias for int32

	// Boolean
	var myBool bool = true

	// Print all types
	fmt.Printf("Int: %d (type: %T)\n", myInt, myInt)
	fmt.Printf("Int8: %d (type: %T)\n", myInt8, myInt8)
	fmt.Printf("Int16: %d (type: %T)\n", myInt16, myInt16)
	fmt.Printf("Int32: %d (type: %T)\n", myInt32, myInt32)
	fmt.Printf("Int64: %d (type: %T)\n", myInt64, myInt64)

	fmt.Printf("Uint: %d (type: %T)\n", myUint, myUint)
	fmt.Printf("Uint8: %d (type: %T)\n", myUint8, myUint8)
	fmt.Printf("Byte: %d (type: %T)\n", myByte, myByte)

	fmt.Printf("Float32: %.2f (type: %T)\n", myFloat32, myFloat32)
	fmt.Printf("Float64: %.15f (type: %T)\n", myFloat64, myFloat64)

	fmt.Printf("Complex64: %g (type: %T)\n", myComplex64, myComplex64)
	fmt.Printf("Complex128: %g (type: %T)\n", myComplex128, myComplex128)

	fmt.Printf("String: %s (type: %T)\n", myString, myString)
	fmt.Printf("Rune: %c (type: %T)\n", myRune, myRune)
	fmt.Printf("Boolean: %t (type: %T)\n", myBool, myBool)
}
