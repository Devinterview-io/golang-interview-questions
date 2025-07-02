// 03_zero_values.go - Demonstrates zero values in Go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("=== GO ZERO VALUES DEMONSTRATION ===")

	// Basic types zero values
	var (
		b  bool
		i  int
		f  float64
		c  complex128
		s  string
		r  rune
		bt byte
	)

	fmt.Printf("bool zero value: %t\n", b)
	fmt.Printf("int zero value: %d\n", i)
	fmt.Printf("float64 zero value: %f\n", f)
	fmt.Printf("complex128 zero value: %g\n", c)
	fmt.Printf("string zero value: '%s' (empty string)\n", s)
	fmt.Printf("rune zero value: %d\n", r)
	fmt.Printf("byte zero value: %d\n", bt)

	// Composite types zero values
	var (
		slice  []int
		dmap   map[string]int
		ptr    *int
		iface  interface{}
		ch     chan int
		fn     func()
		person Person
	)

	fmt.Printf("\nslice zero value: %v (nil: %t)\n", slice, slice == nil)
	fmt.Printf("map zero value: %v (nil: %t)\n", dmap, dmap == nil)
	fmt.Printf("pointer zero value: %v (nil: %t)\n", ptr, ptr == nil)
	fmt.Printf("interface zero value: %v (nil: %t)\n", iface, iface == nil)
	fmt.Printf("channel zero value: %v (nil: %t)\n", ch, ch == nil)
	fmt.Printf("function zero value: %v (nil: %t)\n", fn, fn == nil)
	fmt.Printf("struct zero value: %+v\n", person)
}
