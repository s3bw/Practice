package main

import "fmt"

// Function takes two integers and returns an int
func plus(a int, b int) int {
	return a + b
}

// Consecutive parameters of the same type
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// Calling function
	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res)
}
