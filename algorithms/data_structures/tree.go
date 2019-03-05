package main

import "fmt"

type tree struct {
	item   int
	parent *tree
	left   *tree
	right  *tree
}

// CreateTree needs to return only the root
func CreateTree() *tree {
	// I've hard coded the tree from pg.81 here
	/* 	2
	   / \
	  1   7
	     / \
		4   8
	   / \
	  3   6
	     /
		5
	*/
	return &tree{
		item: 2,
		left: &tree{item: 1},
		right: &tree{
			item: 7,
			left: &tree{
				item: 4,
				left: &tree{item: 3},
				right: &tree{
					item: 6,
					left: &tree{item: 5},
				},
			},
			right: &tree{item: 8},
		},
	}
}

// search proceeds either left or right depending upon
// whether i occurs before or after the root key.
func search(node *tree, i int) *tree {
	if node == nil {
		return nil
	}
	if node.item == i {
		return node
	}
	if i < node.item {
		return search(node.left, i)
	} else {
		return search(node.right, i)
	}
}

func getMin(node *tree) *tree {
	if node == nil {
		return nil
	}
	min := node
	for min.left != nil {
		min = min.left
	}
	return min
}

func main() {
	root := CreateTree()
	fmt.Println(getMin(root).item)
	fmt.Println(search(root, 4).item)
}
