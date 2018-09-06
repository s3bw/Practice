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

import "math"

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
