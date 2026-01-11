package main

import "fmt"

const (
	// Create a huge constant by shifting a 1 bit left 100 places
	Big = 1 << 100
	// Shift it back down again to get a small number
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 { return x * 0.1 }

func showNumberConstants() {
	fmt.Println("Small as int:", needInt(Small))
	fmt.Println("Small as float64:", needFloat(Small))
	fmt.Println("Big as float64:", needFloat(Big))
	// The following line would cause an error: constant overflows int
	// fmt.Println("Big as int:", needInt(Big))
}
