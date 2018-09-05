package main

import "fmt"
import "./trees"

func b_trees() {
	t1 := trees.New(10, 1)
	// trees.BF(t1)
	trees.Walk(t1)
	gt := trees.Draw(t1)
	fmt.Println(gt.Print())
	fmt.Println("Depth: ", trees.Depth(t1))
}

func main() {
	b_trees()
}
