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

import "tree"
import "b_tree"

func main() {
	t := tree.New(10, 1)
	b_tree.Walk(t)
}
