package main

import "fmt"

type node struct {
	// item contains the data item
	item string
	// next points to the successor
	next *node
}

type list struct {
	// head is the head pointer in the list
	head *node
	// track the length
	length int
}

func CreateList() *list {
	return &list{length: 0}
}

// linkedlist support three basic operations
func search(n *node, i string) *node {
	if n == nil {
		return nil
	}
	if n.item == i {
		return n
	} else {
		return search(n.next, i)
	}
}

func (l *list) insert(i string) {
	p := &node{item: i, next: l.head}
	l.head = p
	l.length++
}

// predecessor is needed because it points to the doomed node,
// so it's next pointer must be changed.
func predecessor(n *node, i string) *node {
	if (n == nil) || (n.next == nil) {
		return nil
	}
	if n.next.item == i {
		return n
	} else {
		return predecessor(n.next, i)
	}
}

func (l *list) delete(i string) {
	doomed := search(l.head, i)
	if doomed != nil {
		pred := predecessor(l.head, i)
		if pred == nil {
			l.head = doomed.next
		} else {
			pred.next = doomed.next
		}
		l.length--
	}
}

func (l *list) print() {
	n := l.head
	fmt.Print("Head -> ")
	for n != nil {
		fmt.Printf("%s -> ", n.item)
		n = n.next
	}
	fmt.Print("nil\n")
}

func _main() {
	// This will result in linkedlist
	// head -> a -> b -> c
	newList := CreateList()
	newList.insert("c")
	newList.insert("b")
	newList.insert("a")

	newList.print()
	newList.delete("b")
	newList.print()
}
