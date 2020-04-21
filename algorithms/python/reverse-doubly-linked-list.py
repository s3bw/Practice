"""
Reverse a Doubly-Linked List:
"""


class Node:
    def __init__(self, value):
        self.next = None
        self.previous = None
        self.value = value


class LinkedList:
    def __init__(self):
        self.head = None
        self.tail = None
        self.length = 0

    def append(self, node):
        if self.length == 0:
            self.head = node
            self.tail = node
        else:
            print(node.value)
            self.tail.next = node
            self.tail = node
        self.length += 1

    def reverse(self):
        node = self.head
        self.tail = node

        prev = None
        while node:
            future = node.next
            node.next = prev
            node.previous = future
            prev = node
            node = future

        self.head = prev

    def __iter__(self):
        node = self.head
        while node:
            yield node
            node = node.next

    def __repr__(self):
        return "<[" + ", ".join([n.value for n in self]) + "]>"


if __name__ == "__main__":
    a = Node("a")
    b = Node("b")
    c = Node("c")
    linked = LinkedList()
    linked.append(a)
    linked.append(b)
    linked.append(c)
    print(linked)

    # Here we have a linked list.
    linked.reverse()
    print(linked)
