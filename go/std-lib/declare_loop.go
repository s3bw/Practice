package main

import "fmt"

func main() {
	i := 0

runLoop:
	for {
		if i == 5 {
			break runLoop
		} else {
			i++
			fmt.Println(i)
		}
		continue runLoop
	}
}

// A break statement is useful for breaking out of a loop
// when inside of a switch statement, for example.
func fromTheDocs() {
	a := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}
	m, n := 4, 6

OuterLoop:
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch a[i][j] {
			case 0:
				state = Error
				break OuterLoop
			case item:
				state = Found
				break OuterLoop
			}
		}
	}
}
