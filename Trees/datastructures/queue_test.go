package datastructures_test

import "testing"

import "github.com/GiantsLoveDeathMetal/Practice/trees/common"
import "github.com/GiantsLoveDeathMetal/Practice/trees/datastructures"

func TestQueueLinkedList(t *testing.T) {
	queue := datastructures.NewQueueLinkedList()
	testQueue(queue, t)
}

func testQueue(queue datastructures.Queue, t *testing.T) {
	helper := common.Test{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	helper.Assert(t, !queue.IsEmpty(), "Queue should not be empty")
	helper.Assert(t, queue.Size() == 3, "Queue length should be 3")
	queue.Dequeue()
	queue.Dequeue()
	item := queue.Dequeue()
	helper.Assert(t, item.(int) == 3, "Value here should be 3.")
}
