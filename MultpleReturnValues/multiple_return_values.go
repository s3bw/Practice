package main

import "fmt"

// (int, int) shows that this returns two ints
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Multiple assignment
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Subset return values with _
	_, c := vals()
	fmt.Println(c)
}
