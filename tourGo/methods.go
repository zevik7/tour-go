package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Method for Vertex struct
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Method are just functions
func AbsFunction(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Method continued: declare a method on non-struct type
type MyFloat float64

func (f MyFloat) AbsNonStruct() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Pointer function
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func demonstrateMethods() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunction(v))

	f := MyFloat(math.Sqrt2)
	fmt.Println(f.AbsNonStruct())

	v.Scale(10)
	fmt.Println(v.Abs())

	// When call with pointer function, we must
	ScaleFunc(&v, 10)
	// Not ScaleFunc(v ,10)

	// But method accept both value or pointer receiver
	// Choosing pointer receiver save memory by avoiding copying value for each method call
	v.Scale(5)
	fmt.Println("v.Scale(5)", v.Abs())
	(&v).Scale(5)
	fmt.Println("(&v).Scale(5)", v.Abs())
}
