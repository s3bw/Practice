package main

// The primary mechanism for managing state in Go is communication over channels
// There are a few other options for managing state though. Here we'll look at
// using the sync/atomic package for atomic counters accessed by multiple
// gorountines.

import "fmt"
import "time"
import "sync/atomic"

func main() {

	// We'll use an unsigned integer to represent out (always-positive) counter.
	var ops uint64

	for i := 0; i < 50; i++ {
		go func() {
			for {

				// To atomically increment the counter we use AddUint64, giving
				// it the memory address of our ops counter with the & syntax.
				atomic.AddUint64(&ops, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
