package main

import "fmt"

// Typed collections of fields
// useful for grouping data together
type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(person{"Bob", 20})

	// You can init by name
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields are zero valued
	fmt.Println(person{name: "Fred"})

	// yields a pointer to a struct
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	// Access struct fields with a dot
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
