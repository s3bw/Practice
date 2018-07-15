package main

import "fmt"

func main() {

	// Creating an array which holds 5 ints
	// this is zero valued i.e. an array containing
	// only 0.
	var a [5]int
	fmt.Println("empty:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Like python len
	fmt.Println("len:", len(a))

	// Declare and init an array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declare:", b)

	// Compose multi-dimensional array
	var twoD [2][3]int
	// init condition after
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
