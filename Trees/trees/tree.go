package trees

import "time"
import "math"
import "math/rand"

type Tree struct {
	Left   *Tree
	Value  int
	Right  *Tree
	height float64
	root   bool
	parent *Tree
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

func IsBalanced(t *Tree) bool {
	var lh, rh float64

	if t == nil {
		return true
	}
	lh = Depth(t.Left)
	rh = Depth(t.Right)
	if (math.Abs(lh-rh) <= 1) && IsBalanced(t.Left) && IsBalanced(t.Right) {
		return true
	}
	return false
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
