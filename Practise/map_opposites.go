package main

import "fmt"

const (
	RIGHT direction = 1 + iota
	LEFT
	UP
	DOWN
	NO_MOVE
)

type direction int

func main() {
	m := make(map[direction]direction)

	// Set key pairs
	m[LEFT] = RIGHT
	m[RIGHT] = LEFT
	m[UP] = DOWN
	m[DOWN] = UP
	m[NO_MOVE] = NO_MOVE

	pattern := []direction{LEFT, LEFT, LEFT}

	fmt.Println(pattern)
	for i := 0; i < len(pattern); i++ {
		pattern[i] = m[pattern[i]]
	}
	fmt.Println(pattern)
}
