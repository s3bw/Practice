package main

import "fmt"
import "time"

// Go select lets you wait on multiple channel operations
// combining goroutines and channels with select is a powerful
// feature of go.
func main() {

	// Here we select between two channels
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We use select to await both of these values simultaneously
	// printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
		// Note that the total execution time is only ~2 seconds since
		// both the 1 and 2 second sleeps execute concurrently.
	}
}
