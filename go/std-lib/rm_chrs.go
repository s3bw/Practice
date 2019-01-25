package main

import "fmt"

func main() {
	// Note: here '(' matters as well as the double quotations '"'
	line := []byte("All Done!?")

	// Remove 2nd last character from line, if it's an '!'
	if n := len(line); n >= 2 && line[n-2] == '!' && line[n-1] == '?' {
		line[n-2] = '?'
		line = line[:n-1]
	}
	// Prints: '[65 108 ... 101 63]'
	fmt.Println(line)

	// Prints: 'All Done?'
	fmt.Println(string(line))

	// Prints: '101'
	fmt.Println('e')

	// Print: 'e'
	fmt.Println("e")

	fmt.Println(`e`)
}
