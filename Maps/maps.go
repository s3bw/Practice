package main

import "fmt"

func main() {

	// Basically dictionaries or hashes
	m := make(map[string]int)

	// Set key pairs
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)
	// Outputs:
	// map: map[k1:7 k2:13]

	// Assign v1 to the value in dict by key
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// Prints number of pairs
	fmt.Println("len:", len(m))

	// Removes key from map
	delete(m, "k2")
	fmt.Println("map:", m)

	// Maps/maps.go:29:13: cannot use _ as value
	// This returns a boolean to disambiguate between
	// missing keys and keys with zero values like 0
	_, prs := m["k1"]
	fmt.Println("exists:", prs)

	// Declare and init
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
