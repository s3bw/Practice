// In computer science, a binary tree is a tree
// data structure in which each node has at most
// two children, which are referred to as the left
// child and the right child.
package main

import "fmt"
import "math"
import "strconv"
import "math/rand"

import "github.com/disiqueira/gotree"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Trees may be of different shapes, but have the
// same contents. For example:
//
// 		4                 6
// 	  2   6            4     7
//   1 3 5 7         2   5
//                  1 3

// Walk traverses a tree; depth-first
func Walk(t *Tree) {
	if t == nil {
		return
	}
	Walk(t.Left)
	fmt.Println(t.Value)
	Walk(t.Right)
}

// Determine the depth of the tree
func Depth(t *Tree) float64 {
	var dleft, dright float64
	if t == nil {
		return 0
	}
	if t.Left == nil {
		dleft = 0
	} else {
		dleft = Depth(t.Left)
	}
	if t.Right == nil {
		dright = 0
	} else {
		dright = Depth(t.Right)
	}
	return math.Max(dleft, dright) + 1
}

// Create a new binary tree
func New(n, k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (v + 1))
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

// Drawing tree
// I Need to BFS this with a list see
// https://github.com/maximelamure/algorithms/blob/master/datastructure/queuelinkedlist.go
func Draw(t *Tree, go_tree gotree.Tree) {
	if t.Left != nil {
		tleft := go_tree.Add(strconv.Itoa(t.Left.Value))
		Draw(t.Left, tleft)
	}
	fmt.Println(t.Value)
	if t.Right != nil {
		tright := go_tree.Add(strconv.Itoa(t.Right.Value))
		Draw(t.Right, tright)
	}
}

func main() {
	t1 := New(30, 1)
	go_tree := gotree.New(strconv.Itoa(t1.Value))
	fmt.Println(t1.Value)
	Draw(t1, go_tree)
	fmt.Println(go_tree.Print())
	Walk(t1)
	fmt.Println("Depth: ", Depth(t1))
}
