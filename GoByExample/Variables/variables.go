package main

import "fmt"

func main() {

	// Declaring a var
	var a = "initial"
	fmt.Println(a)

	// Multi Declare and assigning types
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Declare boolean
	var d = true
	fmt.Println(d)

	// Declare without assigning value are given
	// zero value, e.g. int = 0
	var e int
	fmt.Println(e)

	// Short hand for declare and init a var
	// e,g, for f string = "short"
	f := "short"
	fmt.Println(f)
}
