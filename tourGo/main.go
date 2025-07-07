package main

import (
	"fmt"
	"math"
	"runtime"
	"strings"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(math.Pi)
	fmt.Println(Add(1, 2))
	fmt.Println(swap("Phu", "Minh"))
	fmt.Println(Add(1, 2))
	fmt.Println(split(17))

	// Create a var
	var i, j int = 1, 2
	y, z := 3, 4

	fmt.Println(i, j, y, z)

	// Test non-initializer
	var kk complex128

	fmt.Println(kk)

	// Test type conversions
	var keke int = 42
	f := keke

	fmt.Println(f)

	// Test for loop
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	// Optional: init and post statement can be blanked, so for is like "while" in Go
	hehe := 1

	for hehe < 1000 {
		hehe += hehe
	}

	fmt.Println(sum)

	// If statement with init part
	if k := 1 + 1; k < 1 {
		fmt.Println(k)
	} else {
		fmt.Println("K: ", k+1)
	}

	darwin := "darwin"
	linux := "linux"

	switch os := runtime.GOOS; os {
	case darwin:
		fmt.Println("MacOS")
	case linux:
		fmt.Println("Linux")
	default:
		fmt.Printf("%s. \n", os)
	}

	// Pointer in Go
	var p *int

	ii := 42
	p = &ii

	fmt.Println(*p)
	fmt.Println(p)

	// Structs in Go
	type Vertex struct {
		X int
		Y int
	}

	fmt.Println(Vertex{1, 2})
	fmt.Println(Vertex{1, 2}.X)

	vv := Vertex{3, 4}
	pp := &vv
	pp.X = 1e9

	fmt.Println(vv)

	var (
		v1  = Vertex{1, 2}
		v2  = Vertex{X: 1}
		v3  = Vertex{}
		ppp = &Vertex{1, 2}
	)

	fmt.Println(v1, v2, v3, ppp)

	// Array in go
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	fmt.Println(arr)

	// Slices in go
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var slicee []int = primes[1:4]
	fmt.Println(slicee)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b)
	s := make([]int, 3, 5) // Create a slice with length 3 and capacity 5
	fmt.Printf("Initial slice: len=%d, cap=%d, value=%v\n", len(s), cap(s), s)

	s = append(s, 10) // Append an element
	fmt.Printf("After append: len=%d, cap=%d, value=%v\n", len(s), cap(s), s)

	s = append(s, 20, 30) // Append more elements
	fmt.Printf("After more appends: len=%d, cap=%d, value=%v\n", len(s), cap(s), s)

	b[0] = "XXXXXX"

	fmt.Println(a, b)

	fmt.Println(names)

	// Slice defaults

	ss := []int{2, 3, 5, 7, 11, 13}

	ss = ss[1:4]

	fmt.Println(ss)

	ss = ss[1:]
	fmt.Println(ss)

	// Slice length and capacity

	sss := []int{2, 3, 5, 7, 11, 13}
	printSlice(sss)

	sss = sss[:0]
	printSlice(sss)

	sss = sss[:4]
	printSlice(sss)

	sss = sss[2:]
	printSlice(sss)

	ssss := make([]int, 3, 5) // Create a slice with length 3 and capacity 5
	fmt.Printf("Initial slice: len=%d, cap=%d, value=%v\n", len(ssss), cap(ssss), ssss)

	ssss = append(ssss, 10) // Append an element
	fmt.Printf("After append: len=%d, cap=%d, value=%v\n", len(ssss), cap(ssss), ssss)

	ssss = append(ssss, 20, 30) // Append more elements
	fmt.Printf("After more appends: len=%d, cap=%d, value=%v\n", len(ssss), cap(ssss), ssss)

	// Slices of slices

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println(board)

	var testAppend []int
	printSlice(testAppend)

	testAppend = append(testAppend, 0)
	printSlice(testAppend)

	testAppend = append(testAppend, 1)
	printSlice(testAppend)

	testAppend = append(testAppend, 2, 3, 4, 5)
	printSlice(testAppend)

	var pow = []int{1, 2, 3, 4, 5, 6, 7}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// omit value in range
	for _, value := range pow {
		fmt.Printf("%d", value)
	}

	// Draw picture with vector
	Pic(5, 5)

	// Map
	var mmm map[string]Vertex
	mmm = make(map[string]Vertex)

	mmm["Zevik"] = Vertex{1, 2}
	fmt.Println(mmm["Zevik"])

	// Map literals
	mmm = map[string]Vertex{
		"Zevik": {1, 2},
		"John":  {3, 4},
	}
	fmt.Println(mmm["John"])

	// Mutating map
	m := make(map[string]int)

	m["Zevik"] = 1
	m["John"] = 2

	delete(m, "Zevik")

	v, present := m["Zevik"]
	fmt.Println(v, present)

	// Function values

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// Function closures
	pos, neg := adder(), adder()

	for i := 0; i < 5; i++ {
		fmt.Println(pos(i), neg(-i))
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		result[y] = make([]uint8, dx)

		for x := 0; x < dx; x++ {
			result[y][x] = uint8((x ^ y))
		}
	}

	return result
}

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		counts[word]++
	}

	return counts
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
