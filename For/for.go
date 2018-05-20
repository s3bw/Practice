package main

import "fmt"

func main() {

	// Shorthand
	i := 1
	for i <= 3 {
		fmt.Println(i)

		// In Go as in Python
		i += 1
	}

	// Not so Python init/condition/after loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		// In Go as in Python
		break
	}

	// Continue to next iter
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
