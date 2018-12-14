// Package trees containing DSW algorithm
//
// Day-Stout-Warren algorithm works in two parts:
//
// 	1] Transform the tree into a linked list
// 	   using an in-order traversal.
//  2] Using a series of left rotations
// 	   balance the tree.
package trees

import "fmt"

// VineTransformation transform tree into linked-list
func VineTransformation(t *Tree) *Tree {
	var gp, leftChild, parent *Tree

	parent = t
	for parent != nil {
		leftChild = parent.Left
		if leftChild != nil {
			t = rotateRight(t, gp, parent, leftChild)
			parent = leftChild
		} else {
			gp = parent
			parent = parent.Right
		}
	}
	return t
}

func rotateRight(t, gp, parent, leftChild *Tree) *Tree {
	if gp != nil {
		gp.Right = leftChild
	} else {
		t.root = false
		leftChild.root = true
		t = leftChild
	}

	parent.Left = leftChild.Right
	leftChild.Right = parent
	return t
}

// BalanceTree with left rotations
func BalanceTree(t *Tree) {
	var n float64
	var m int

	n = Depth(t)
	m = greatestPowerOfTwo(uint(n))

	makeRotations(t, int(n)-m)
	for m > 1 {
		m /= 2
		makeRotations(t, m)
	}
	fmt.Println(n)
	fmt.Println(m)
}

func greatestPowerOfTwo(n uint) int {
	return 1 << uint(msb(n))
}

func msb(n uint) int {
	ndx := 0
	for n > 1 {
		n = n >> 1
		ndx++
	}
	return ndx
}

func makeRotations(t *Tree, bound int) {
}
