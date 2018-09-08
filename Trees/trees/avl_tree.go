// A self-balancing binary search tree. In an AVL
// tree, the heights of the two child subtrees of
// any node differ by at most one; if at any time
// they differ by more than one, rebalancing is
// done to restore this property.

// AVL trees are often compared with red-black
// trees because both support the same set of
// operations. For lookup-instensive applications
// AVL trees are faster than red-black trees
// because they are more strictly balanced.

package trees

import "time"
import "math/rand"

func FindRoot(t *Tree) *Tree {
	for t.parent != nil {
		t = t.parent
	}
	return t
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

func Find(t *Tree, k int) *Tree {
	if k == t.Value {
		return t
	} else if k < t.Value {
		if t.Left != nil {
			return Find(t.Left, k)
		}
		return nil
	} else {
		if t.Right != nil {
			return Find(t.Right, k)
		}
		return nil
	}
}

func FindMin(t *Tree) *Tree {
	if t.Left != nil {
		return FindMin(t.Left)
	}
	return t
}

func LeftRotate(t *Tree) {
	var y *Tree

	y = t.Right
	y.parent = t.parent
	if y.parent == nil {
		y.root = true
	} else {
		if *y.parent.Left == *t {
			y.parent.Left = y
		} else if *y.parent.Right == *t {
			y.parent.Right = y
		}
	}

	t.Right = y.Left
	if t.Right != nil {
		t.Right.parent = t
	}
	y.Left = t
	t.parent = y
	UpdateHeight(t)
	UpdateHeight(y)
}

func RightRotate(t *Tree) {
	var y *Tree

	y = t.Left
	y.parent = t.parent
	if y.parent == nil {
		y.root = true
	} else {
		if y.parent.Left == t {
			y.parent.Left = y
		} else if *y.parent.Right == *t {
			y.parent.Right = y
		}
	}
	t.Left = y.Right
	if t.Left != nil {
		t.Left.parent = t
	}
	y.Right = t
	t.parent = y
	UpdateHeight(t)
	UpdateHeight(y)
}

func Rebalance(t *Tree) {
	for t != nil {
		UpdateHeight(t)
		if Depth(t.Left) >= 2+Depth(t.Right) {
			if Depth(t.Left.Left) >= Depth(t.Left.Right) {
				RightRotate(t)
			} else {
				LeftRotate(t.Left)
				RightRotate(t)
			}
		} else if Depth(t.Right) >= 2+Depth(t.Left) {
			if Depth(t.Right.Right) >= Depth(t.Right.Left) {
				LeftRotate(t)
			} else {
				RightRotate(t.Right)
				LeftRotate(t)
			}
		}
		t = t.parent
	}
}
