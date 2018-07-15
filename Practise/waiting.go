package main

import "fmt"
import "time"

// Types of food that can be found
const (
	APPLES food = 1 + iota
	PEARS
	ORANGES
)

type food int

// Gather gets food for the stockpile
func gatherer(stockpile chan food) {
	fmt.Print("finding...")
	time.Sleep(3 * time.Second)
	fmt.Println("found!")
	stockpile <- APPLES
}

func main() {

	stockpile := make(chan food)
	go gatherer(stockpile)

	select {
	// Wait until the gather returns food till it continues
	case f := <-stockpile:
		{
			switch f {
			case APPLES:
				fmt.Println("Got apples!")
			case PEARS:
				fmt.Println("Got pears!")
			case ORANGES:
				fmt.Println("Got oranges!")
			}
		}
	}
}
