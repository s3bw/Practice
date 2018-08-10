package main

import "fmt"

func main() {

	// Sum the contents of a slice
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// These auto enumerate with index value
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Iterate over the key pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Iterate over just the keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over unicode code points
	// the first value is the starting byte index of the
	// rune and the second the rune itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
