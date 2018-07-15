package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// *int means this function takes an int pointer
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i gives the memory address of i.
	// i.e. a pointer to i
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointers can be printed.
	fmt.Println("pointer:", &i)
}
