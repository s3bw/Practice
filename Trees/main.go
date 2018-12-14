package main

import "fmt"
import "./trees"

func stats(t *trees.Tree) {
	gt := trees.Draw(t)
	fmt.Println(gt.Print())
	fmt.Println("Depth: ", trees.Depth(t))
}

func isbalanced(t *trees.Tree) {
	if trees.IsBalanced(t) {
		fmt.Println("Tree is balanced!")
	} else {
		fmt.Println("Tree is NOT balanced!")
	}
}

func b_trees() {
	t1 := trees.New(20, 1)
	// trees.BF(t1)
	// trees.WalkInOrder(t1)
	stats(t1)
	isbalanced(t1)
}

func avl_trees() {
	t := trees.NewAVL(20, 1)
	stats(t)
	x := trees.FindMin(t)
	fmt.Println(x.Value)
	isbalanced(t)
}

func dswAlgo() {
	fmt.Println("Unbalanced tree.")
	t1 := trees.New(10, 1)
	stats(t1)
	fmt.Println("\nVine Transformation.")
	t1 = trees.VineTransformation(t1)
	stats(t1)
	trees.BalanceTree(t1)
	stats(t1)
}

// func rb_trees() {
// 	t := trees.NewRB(20, 1)
// 	// gt := trees.DrawRB(t)
// 	// fmt.Println(gt.Print())
// }

func main() {
	dswAlgo()
	// b_trees()
	// avl_trees()
}
