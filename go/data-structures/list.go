package main

import "fmt"

// List implementation of the abstract data type
type List interface {
	Append(value interface{})
	Size() int
	IsEmpty() bool
	Insert(i int, value interface{}) error
	Pop(i int) (interface{}, error)
	Get(i int) (interface{}, error)
	Iterate() <-chan interface{}
}

type list struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	prev  *node
	next  *node
	value interface{}
}

// CreateList will return an implementation of a linked list
func CreateList() List {
	return &list{}
}

func (l *list) Iterate() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for i := 0; i < l.Size(); i++ {
			if l.IsEmpty() {
				break
			}
			value, err := l.Get(i)
			if err != nil {
				break
			}
			ch <- value
		}
		close(ch)
	}()
	return ch
}

func (l *list) Get(i int) (interface{}, error) {
	if i < 0 || i > l.Size() {
		return nil, fmt.Errorf("Index out of bounds")
	}
	node := l.head
	j := 0
	for j < i-1 {
		j++
		node = node.next
	}
	return &node.value, nil
}

func (l *list) Pop(i int) (interface{}, error) {
	if i < 0 || i > l.Size() {
		return nil, fmt.Errorf("Index out of bounds")
	}
	node := l.head
	j := 0
	for j < i-1 {
		j++
		node = node.next
	}
	remove := node.next
	node.next = remove.next
	l.length--
	return &remove.value, nil
}

func (l *list) Insert(i int, value interface{}) error {
	if i < 0 || i > l.Size() {
		return fmt.Errorf("Index out of bounds")
	}
	addNode := node{value: value}
	if i == 0 {
		addNode.next = l.head
		l.head = &addNode
		return nil
	}
	node := l.head
	j := 0
	for j < i-2 {
		j++
		node = node.next
	}
	addNode.next = node.next
	node.next = &addNode
	l.length++
	return nil
}

func (l *list) Append(value interface{}) {
	oldTail := l.tail
	l.tail = &node{}
	l.tail.value = value

	if l.IsEmpty() {
		l.head = l.tail
	} else {
		oldTail.next = l.tail
	}
	l.length++
}

func (l *list) IsEmpty() bool {
	return l.Size() == 0
}

func (l *list) Size() int {
	return l.length
}
