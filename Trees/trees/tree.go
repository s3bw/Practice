package trees

import "time"
import "math/rand"

type Tree struct {
	Left   *Tree
	Value  int
	Right  *Tree
	height float64
	root   bool
	parent *Tree
}

// Create a new binary tree
func New(n, k int) *Tree {
	var t *Tree

	// Seed random b tree
	rand.Seed(time.Now().UnixNano())
	for _, v := range rand.Perm(n) {
		t = Insert(t, (v + 1))
	}
	return t
}

func Insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{
			nil,
			v,
			nil,
			0,
			false,
			nil}
	}
	if v < t.Value {
		t.Left = Insert(t.Left, v)
		return t
	}
	t.Right = Insert(t.Right, v)
	return t
}

func NewNode(k int) *Tree {
	return &Tree{
		nil,
		k,
		nil,
		0,
		false,
		nil}
}

func NewAVL(n, k int) *Tree {
	var t, root *Tree

	// Seed random b tree
	rand.Seed(time.Now().UnixNano())
	for _, v := range rand.Perm(n) {
		t = NewNode(v + 1)
		if root == nil {
			root = t
			root.root = true
		} else {
			InsertAVL(root, t)
		}
		Rebalance(t)
		root = FindRoot(t)
	}
	return root
}

func InsertAVL(a, n *Tree) {
	if n == nil {
		return
	}
	if n.Value < a.Value {
		if a.Left == nil {
			n.parent = a
			a.Left = n
		} else {
			InsertAVL(a.Left, n)
		}
	} else {
		if a.Right == nil {
			n.parent = a
			a.Right = n
		} else {
			InsertAVL(a.Right, n)
		}
	}
	return
}

func FindRoot(t *Tree) *Tree {
	for t.parent != nil {
		t = t.parent
	}
	return t
}

func Successor(t *Tree) *Tree {
	var current *Tree

	if t.Right != nil {
		return FindMin(t.Right)
	}
	current = t
	for current.parent != nil && current == current.parent.Right {
		current = current.parent
	}
	return current.parent
}
