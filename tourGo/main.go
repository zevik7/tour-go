package main

import (
	"fmt"
	"math"
	"runtime"
	"strings"
)

// =============================================================================
// MAIN FUNCTION - Entry point demonstrating Go language fundamentals
// =============================================================================

func main() {
	fmt.Println("\n========== BASIC FUNCTIONS ==========")
	demonstrateBasicFunctions()

	fmt.Println("\n========== VARIABLES & TYPES ==========")
	demonstrateVariablesAndTypes()

	fmt.Println("\n========== CONTROL FLOW ==========")
	demonstrateControlFlow()

	fmt.Println("\n========== POINTERS ==========")
	demonstratePointers()

	fmt.Println("\n========== STRUCTS ==========")
	demonstrateStructs()

	fmt.Println("\n========== ARRAYS ==========")
	demonstrateArrays()

	fmt.Println("\n========== SLICES ==========")
	demonstrateSlices()

	fmt.Println("\n========== MAPS ==========")
	demonstrateMaps()

	fmt.Println("\n========== FUNCTION VALUES & CLOSURES ==========")
	demonstrateFunctionValues()

	fmt.Println("\n========== Methods =============================")
	demonstrateMethods()
}

// =============================================================================
// SECTION 1: BASIC FUNCTIONS
// Demonstrates: function calls, multiple return values, named returns
// =============================================================================

func demonstrateBasicFunctions() {
	// Using math package constant
	fmt.Println("Pi from math package:", math.Pi)

	// Simple function with two parameters (defined in add.go)
	fmt.Println("Add(1, 2) =", Add(1, 2))

	// Function with multiple return values (defined in swap.go)
	first, second := swap("Phu", "Minh")
	fmt.Printf("swap('Phu', 'Minh') = %s, %s\n", first, second)

	// Function with named return values (naked return)
	x, y := split(17)
	fmt.Printf("split(17) = %d, %d\n", x, y)
}

// =============================================================================
// SECTION 2: VARIABLES AND TYPES
// Demonstrates: variable declarations, type inference, zero values, type conversions
// =============================================================================

func demonstrateVariablesAndTypes() {
	// Variable declaration with explicit type and initialization
	var i, j int = 1, 2

	// Short variable declaration with type inference
	y, z := 3, 4

	// Variables with initializers
	var k, l = 5, 6

	fmt.Println("Multiple variables:", i, j, y, z, k, l)

	// Variables without initializers get zero values (0, false, "", etc.)
	var kk complex128
	fmt.Println("Zero value of complex128:", kk)

	// Show basic types
	showTypes()
	showZeroValues()

	// Type conversions - Go requires explicit type conversion
	var original int = 42
	converted := original // Same type, no conversion needed
	fmt.Println("Converted value:", converted)

	const testConstantPi = 3.14
	var floatValue float64 = testConstantPi // Implicit conversion from constant
	fmt.Println("Constant to float64:", floatValue)

	// Show number constants
	showNumberConstants()
}

// =============================================================================
// SECTION 3: CONTROL FLOW
// Demonstrates: for loops, if statements with init, switch statements
// =============================================================================

func demonstrateControlFlow() {
	// Standard for loop with init, condition, and post statement
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum from 0 to 9:", sum)

	// For loop with range
	sum2 := 0
	for _, value := range []int{1, 2, 3, 4, 5} {
		sum2 += value
	}

	// For loop as a "while" loop (init and post statements are optional)
	counter := 1
	for counter < 1000 {
		counter += counter
	}
	fmt.Println("Counter doubled until >= 1000:", counter)

	// If statement with initialization (scope of k is limited to if-else block)
	if k := 1 + 1; k < 1 {
		fmt.Println("k is less than 1:", k)
	} else {
		fmt.Println("k is not less than 1, k+1 =", k+1)
	}

	// Switch statement with initialization
	darwin := "darwin"
	linux := "linux"

	switch os := runtime.GOOS; os {
	case darwin:
		fmt.Println("Operating System: MacOS")
	case linux:
		fmt.Println("Operating System: Linux")
	default:
		fmt.Printf("Operating System: %s\n", os)
	}
}

// =============================================================================
// SECTION 4: POINTERS
// Demonstrates: pointer declaration, address operator &, dereference operator *
// =============================================================================

func demonstratePointers() {
	// Declare a pointer to an int
	var p *int

	// Create a variable and get its address
	ii := 42
	p = &ii // & generates a pointer to its operand

	// Dereference the pointer to access the value
	fmt.Println("Value at pointer:", *p)
	fmt.Println("Memory address:", p)

	// Modify value through pointer
	*p = 21
	fmt.Println("Updated value:", ii) // ii is now 21
}

// =============================================================================
// SECTION 5: STRUCTS
// Demonstrates: struct definition, initialization, field access, pointers to structs
// =============================================================================

func demonstrateStructs() {
	// Vertex represents a 2D point with X and Y coordinates
	type Vertex struct {
		X int
		Y int
	}

	// Creating struct instances
	fmt.Println("Create struct with values:", Vertex{1, 2})
	fmt.Println("Access struct field:", Vertex{1, 2}.X)

	// Accessing struct fields through pointers
	vv := Vertex{3, 4}
	pp := &vv
	pp.X = 1e9 // Go allows direct field access through pointer (no need for (*pp).X)
	fmt.Println("Modified struct through pointer:", vv)

	// Different ways to initialize structs
	var (
		v1  = Vertex{1, 2}  // Full initialization
		v2  = Vertex{X: 1}  // Named fields (Y gets zero value)
		v3  = Vertex{}      // All fields get zero values
		ppp = &Vertex{1, 2} // Creates struct and returns pointer
	)
	fmt.Println("Various struct initializations:", v1, v2, v3, ppp)
}

// =============================================================================
// SECTION 6: ARRAYS
// Demonstrates: array declaration, fixed size, array initialization
// =============================================================================

func demonstrateArrays() {
	// Arrays have fixed size, declared as [n]T where n is size and T is type
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	fmt.Println("Array with 5 elements:", arr)
	fmt.Println("Array length:", len(arr))
}

// =============================================================================
// SECTION 7: SLICES
// Demonstrates: slice creation, slicing operations, length vs capacity,
//               make(), append(), slice references, 2D slices
// =============================================================================

func demonstrateSlices() {
	// Slice literals (dynamically-sized, flexible view into arrays)
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("Slice of primes:", primes)

	// Slicing creates a view into the original array
	var slice []int = primes[1:4] // Elements at index 1, 2, 3
	fmt.Println("Slice [1:4]:", slice)

	// Slices reference underlying arrays - modifications are visible
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println("Original array:", names)

	a := names[0:2] // [John, Paul]
	b := names[1:3] // [Paul, George]
	fmt.Println("Two slices from same array:", a, b)

	b[0] = "XXXXXX" // Modifies underlying array
	fmt.Println("After modification:", a, b)
	fmt.Println("Original array modified:", names)

	//The length of a slice is the number of elements it contains.
	//The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	testLengthCap := []int{2, 3, 5, 7, 11, 13}
	printSlice(testLengthCap)

	fmt.Println("Slice a slice to make length zero:")
	testLengthCap = testLengthCap[:0]
	printSlice(testLengthCap)

	fmt.Println("Extend length of slice:")
	testLengthCap = testLengthCap[:4] // Extend its length.
	printSlice(testLengthCap)

	fmt.Println("Drop first two values of slice, the capacity decreases:")
	testLengthCap = testLengthCap[2:] // Drop its first two values.
	printSlice(testLengthCap)

	// Nil slices
	var nilSlice []int
	fmt.Println("Nil slice:", nilSlice)

	// Creating slices with make([]T, length, capacity)
	s := make([]int, 3, 5) // Length 3, capacity 5
	fmt.Printf("Initial slice: len=%d, cap=%d, value=%v\n", len(s), cap(s), s)

	s = append(s, 10) // Append single element
	fmt.Printf("After append: len=%d, cap=%d, value=%v\n", len(s), cap(s), s)

	s = append(s, 20, 30) // Append multiple elements
	fmt.Printf("After more appends: len=%d, cap=%d, value=%v\n", len(s), cap(s), s)

	// Slice defaults: omit high or low bounds to use their defaults
	ss := []int{2, 3, 5, 7, 11, 13}
	ss = ss[1:4] // [3 5 7]
	fmt.Println("Slice [1:4]:", ss)

	ss = ss[1:] // From index 1 to end
	fmt.Println("Slice [1:]:", ss)

	// Appending to nil slices (nil slice has length and capacity 0)
	var testAppend []int
	printSlice(testAppend)

	testAppend = append(testAppend, 0)
	printSlice(testAppend)

	testAppend = append(testAppend, 1)
	printSlice(testAppend)

	testAppend = append(testAppend, 2, 3, 4, 5)
	printSlice(testAppend)

	// Slices of slices (2D slices/matrices)
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	fmt.Println("2D slice (board):", board)

	// Range loops iterate over slice with index and value
	pow := []int{1, 2, 4, 8, 16, 32, 64}
	for i, v := range pow {
		fmt.Printf("Index %d has value %d\n", i, v)
	}

	// Omit index with _ (blank identifier)
	fmt.Print("Values only: ")
	for _, value := range pow {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

	// Generate a 2D slice using Pic function
	result := Pic(5, 5)
	fmt.Println("Generated 5x5 picture:", result)
}

// =============================================================================
// SECTION 8: MAPS
// Demonstrates: map creation, map literals, insertion, deletion, lookup with "comma ok"
// =============================================================================

func demonstrateMaps() {
	// Maps are key-value pairs, must be created with make
	type Vertex struct {
		X int
		Y int
	}
	var demoMap map[string]Vertex = make(map[string]Vertex)

	demoMap["Zevik"] = Vertex{1, 2}
	fmt.Println("Map value for 'Zevik':", demoMap["Zevik"])

	// Map literals (initialize with values)
	demoMap = map[string]Vertex{
		"Zevik": {1, 2},
		"John":  {3, 4},
	}
	fmt.Println("Map value for 'John':", demoMap["John"])

	// Mutating maps: insert, update, delete, test for key presence
	mutatingMap := make(map[string]int)

	mutatingMap["Zevik"] = 1
	mutatingMap["John"] = 2
	fmt.Println("Map with 2 entries:", mutatingMap)

	delete(mutatingMap, "Zevik") // Delete key from map
	fmt.Println("After deleting 'Zevik':", mutatingMap)

	// Test for key presence with two-value assignment
	v, present := mutatingMap["Zevik"]
	fmt.Printf("Key 'Zevik' - value: %d, present: %v\n", v, present)

	// Test WordCount function
	sentence := "hello world hello go"
	counts := WordCount(sentence)
	fmt.Println("Word counts:", counts)
}

// =============================================================================
// SECTION 9: FUNCTION VALUES & CLOSURES
// Demonstrates: functions as values, higher-order functions, closures
// =============================================================================

func demonstrateFunctionValues() {
	// Functions are values too - can be passed around
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println("hypot(5, 12) =", hypot(5, 12))

	// Pass function as argument to another function
	fmt.Println("compute(hypot) =", compute(hypot))
	fmt.Println("compute(math.Pow) =", compute(math.Pow))

	// Function closures - functions that reference variables from outside their body
	pos, neg := adder(), adder() // Two independent closure instances

	for i := range 5 {
		fmt.Printf("pos(%d) = %d, neg(%d) = %d\n", i, pos(i), -i, neg(-i))
	}

	// Fibonacci closure demonstration
	fib := fibonacci()
	fmt.Print("Fibonacci sequence: ")
	for range 10 {
		fmt.Printf("%d ", fib())
	}
	fmt.Println()
}

// =============================================================================
// HELPER FUNCTIONS & UTILITIES
// =============================================================================

// printSlice prints length, capacity, and contents of an int slice
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Pic generates a 2D slice of uint8 using XOR operation
// Can be used to create interesting visual patterns
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)

	for y := range result {
		result[y] = make([]uint8, dx)

		for x := range result[y] {
			result[y][x] = uint8((x ^ y)) // XOR creates patterns
		}
	}

	return result
}

// WordCount returns a map of word counts in a given string
func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	words := strings.FieldsSeq(s) // Split string into words

	for word := range words {
		counts[word]++ // Increment count for each word
	}

	return counts
}

// compute takes a function and applies it to fixed arguments
// Demonstrates passing functions as parameters
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// adder returns a closure that accumulates a sum
// Each closure maintains its own independent sum variable
func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci returns a closure that generates fibonacci numbers
// The closure maintains state across calls
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
