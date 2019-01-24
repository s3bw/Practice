/* Implementing BFS and a Queue*/

package graph

import "sync"

// NodeQueue the queue of Nodes
type NodeQueue struct {
	items []Node
	lock  sync.RWMutex
}

// New creates a new NodeQueue
func (s *NodeQueue) New() *NodeQueue {
	s.lock.Lock()
	s.items = []Node{}
	s.lock.Unlock()
	return s
}

// Enqueue adds a Node to the end of the queue
func (s *NodeQueue) Enqueue(t Node) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// Dequeue removes a Node from the start of the queue
func (s *NodeQueue) Dequeue() *Node {
	var item Node

	s.lock.Lock()
	item, s.items = s.items[0], s.items[1:]
	s.lock.Unlock()
	return &item
}

// Peak returns item next in queue, without removing it
func (s *NodeQueue) Peak() *Node {
	s.lock.RLock()
	item := s.items[0]
	s.lock.RUnlock()
	return &item
}

// IsEmpty returns true if the queue is empty
func (s *NodeQueue) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items) == 0
}

// Size returns the number of Nodes in the queue
func (s *NodeQueue) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}

// BFTraversal goes through nodes in graph BF
func (g *ItemGraph) BFTraversal(f func(*Node)) {
	g.lock.RLock()
	q := NodeQueue{}
	q.New()
	n := g.nodes[0]
	q.Enqueue(*n)
	visited := make(map[*Node]bool)
	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue()
		visited[node] = true
		near := g.edges[*node]

		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				q.Enqueue(*j)
				visited[j] = true
			}
		}
		if f != nil {
			f(node)
		}
	}
	g.lock.RUnlock()
}
