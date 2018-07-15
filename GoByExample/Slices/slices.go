package main

import "fmt"

func main() {

	// builtin make:
	// Slice of empty strings
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// Prints "len: 3"
	fmt.Println("len:", len(s))

	// Slice support several builtin operations
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	// Slices can be copied
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy", c)

	// Slices and slice operations
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// Two dimensional slices
	// blog: https://blog.golang.org/go-slices-usage-and-internals
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		// Inner slices can vary unlike arrays
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
