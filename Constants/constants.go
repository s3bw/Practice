package main

import "fmt"
import "math"

// Declaring a constant value
const s string = "constant"

func main() {
	fmt.Println(s)

	// Constants are basically not variable
	const n = 5000000

	// arithmetic notation
	const d = 3e20 / n
	fmt.Println(d)

	// Numeric constants have to type until
	// explicitly cast.
	// http://hcc-cs.weebly.com/casting.html
	fmt.Println(int64(d))

	// Numbers are dynamically typed in situations
	// that require a type.
	// This function expects float64
	fmt.Println(math.Sin(n))
}
