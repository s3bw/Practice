package main

import "fmt"
import "time"

// Rate limiting is an important mechanism for controlling resource
// utilization and maintaining quality of service. Go elegantly supports
// rate limiting with goroutines, channels and tickers.
func main() {

	// Basic:
	// Suppose we want to limit our handling of incoming request. We'll
	// serve these requests off a channel of the same name.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This limiter channel will receive a value every 200 milliseconds.
	// This is the regulator in our rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on a receive from the limiter channel before serving
	// each request, we limit ourselves to 1 request every 200 milliseconds
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Short bursts of requests can be accomplished by buffering our
	// limiter channel. This burstyLimiter channel will allow bursts
	// of up to 3 events
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Add a new value every 200ms, up to its limit of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Simulate 5 more incoming requests
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("bursty-request", req, time.Now())
	}
}
