package main

import "fmt"

const (
	A command = 1 + iota
	B
	C
	D
)

type command int

// pop first element and append to end of slice
func nextMove(p *[]command) command {
	var move command

	move, *p = (*p)[0], (*p)[1:]
	*p = append(*p, move)
	return move
}

func main() {
	var move command
	pattern := []command{B, A, C}

	// Prints 2
	move = nextMove(&pattern)
	fmt.Println(move)

	// Prints 1
	move = nextMove(&pattern)
	fmt.Println(move)

	// Prints 3
	move = nextMove(&pattern)
	fmt.Println(move)

	// Prints 2
	move = nextMove(&pattern)
	fmt.Println(move)
}
