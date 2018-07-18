package main

import "time"
import "fmt"
import "math/rand"

// Since random ints are deterministic we need to provide
// a seed that changes.
func main() {
	// Create time seed and random number generator
	seedtime := rand.NewSource(time.Now().UnixNano())
	random_generator := rand.New(seedtime)

	// Generate between 0 <= x < 100
	fmt.Println(random_generator.Intn(100))
}
