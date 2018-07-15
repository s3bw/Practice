package main

import "fmt"
import "time"

// Closing a channel indicates that no more values will be sent on it.
// his can be useful to communicate completion to the channel's
// receivers.
func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// the 'more' value will be false if jobs has been closed
			// We use this to notify on done when we've worked all our jobs
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")

				// When we have no more jobs for the worker we'll close
				// the jobs channel
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
		time.Sleep(2 * time.Second)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
