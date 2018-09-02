// In computer science, a binary tree is a tree
// data structure in which each node has at most
// two children, which are referred to as the left
// child and the right child.
package trees

import "fmt"
import "math"
import "strconv"
import "math/rand"

import "github.com/disiqueira/gotree"
import "github.com/GiantsLoveDeathMetal/Practice/trees/datastructures"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Trees may be of different shapes, but have the
// same contents. For example:
//
// 	    4                 6
// 	  2   6            4     7
//   1 3 5 7         2   5
//                  1 3

// Visit provides standard operation to node traversal
func Visit(t *Tree) {
	fmt.Println(t.Value)
}

// Walk traverses a tree; depth-first (In-order)
func Walk(t *Tree) {
	if t == nil {
		return
	}
	Walk(t.Left)
	Visit(t)
	Walk(t.Right)
}

// Breath-first traversal
func BF(t *Tree) {
	queue := datastructures.NewQueueLinkedList()
	queue.Enqueue(t)
	for queue.Size() > 0 {
		n := queue.Dequeue().(*Tree)
		Visit(t)
		if n.Left != nil {
			queue.Enqueue(n.Left)
		}
		if n.Right != nil {
			queue.Enqueue(n.Right)
		}
	}
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

// Breath-first Draw
func Draw(t *Tree) gotree.Tree {
	queue := datastructures.NewQueueLinkedList()
	gqueue := datastructures.NewQueueLinkedList()
	queue.Enqueue(t)
	gtree := gotree.New(strconv.Itoa(t.Value))
	gqueue.Enqueue(gtree)

	for queue.Size() > 0 {
		n := queue.Dequeue().(*Tree)
		gn := gqueue.Dequeue().(gotree.Tree)
		if n.Left != nil {
			queue.Enqueue(n.Left)
			tleft := gn.Add(strconv.Itoa(n.Left.Value))
			gqueue.Enqueue(tleft)
		}
		if n.Right != nil {
			queue.Enqueue(n.Right)
			tright := gn.Add(strconv.Itoa(n.Right.Value))
			gqueue.Enqueue(tright)
		}
	}
	return gtree
}

func main() {
	t1 := New(10, 1)
	// BF(t1)
	Walk(t1)

	gt := Draw(t1)
	fmt.Println(gt.Print())
	fmt.Println("Depth: ", Depth(t1))
}
