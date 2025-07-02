// 12_interfaces.go - Demonstrates interfaces in Go
package main

import (
	"fmt"
	"math"
)

// Define interfaces
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Drawable interface {
	Draw()
}

// Interface composition
type DrawableShape interface {
	Shape
	Drawable
}

// Implement structs
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// Rectangle methods (implements Shape interface)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Draw() {
	fmt.Printf("Drawing a rectangle: %.1fx%.1f\n", r.Width, r.Height)
}

// Circle methods (implements Shape interface)
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Draw() {
	fmt.Printf("Drawing a circle with radius: %.1f\n", c.Radius)
}

// Functions that work with interfaces
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func drawShape(d Drawable) {
	d.Draw()
}

func processDrawableShape(ds DrawableShape) {
	ds.Draw()
	fmt.Printf("Area: %.2f\n", ds.Area())
}

// Empty interface example
func printAnyValue(value interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

func main() {
	fmt.Println("=== INTERFACES DEMONSTRATION ===")

	// Creating instances
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	// Using interfaces
	fmt.Println("--- BASIC INTERFACE USAGE ---")
	printShapeInfo(rect)
	printShapeInfo(circle)

	// Interface variables
	fmt.Println("\n--- INTERFACE VARIABLES ---")
	var shape Shape

	shape = rect
	fmt.Printf("Rectangle - ")
	printShapeInfo(shape)

	shape = circle
	fmt.Printf("Circle - ")
	printShapeInfo(shape)

	// Interface composition
	fmt.Println("\n--- INTERFACE COMPOSITION ---")
	drawShape(rect)
	drawShape(circle)

	processDrawableShape(rect)
	processDrawableShape(circle)

	// Type assertion
	fmt.Println("\n--- TYPE ASSERTION ---")
	shape = rect

	if r, ok := shape.(Rectangle); ok {
		fmt.Printf("It's a rectangle: %+v\n", r)
	}

	if c, ok := shape.(Circle); ok {
		fmt.Printf("It's a circle: %+v\n", c)
	} else {
		fmt.Println("It's not a circle")
	}

	// Type switch
	fmt.Println("\n--- TYPE SWITCH ---")
	shapes := []Shape{rect, circle, Rectangle{Width: 2, Height: 8}}

	for i, s := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		switch v := s.(type) {
		case Rectangle:
			fmt.Printf("Rectangle with width %.1f and height %.1f\n", v.Width, v.Height)
		case Circle:
			fmt.Printf("Circle with radius %.1f\n", v.Radius)
		default:
			fmt.Printf("Unknown shape: %T\n", v)
		}
	}

	// Empty interface
	fmt.Println("\n--- EMPTY INTERFACE ---")
	printAnyValue(42)
	printAnyValue("Hello")
	printAnyValue(3.14)
	printAnyValue(rect)
	printAnyValue([]int{1, 2, 3})

	// Interface with nil
	fmt.Println("\n--- NIL INTERFACES ---")
	var nilShape Shape
	fmt.Printf("Nil interface: %v (nil: %t)\n", nilShape, nilShape == nil)

	var nilPtr *Rectangle
	shape = nilPtr
	fmt.Printf("Interface with nil pointer: %v (nil: %t)\n", shape, shape == nil)

	// Checking for nil safely
	if shape != nil {
		// This would panic if nilPtr methods are called
		fmt.Println("Shape is not nil (but underlying value might be)")
	}
}
