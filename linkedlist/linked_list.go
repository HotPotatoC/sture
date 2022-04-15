package linkedlist

import (
	"errors"
	"fmt"
)

// LinkedList is a linked list.
type LinkedList[V any] struct {
	head  *Node[V]
	tail  *Node[V]
	nSize uint
	cmp   func(V, V) int
}

// NewLinkedList returns a new linked list.
func NewLinkedList[V any](cmp func(V, V) int) *LinkedList[V] {
	return &LinkedList[V]{cmp: cmp}
}

// Append adds a new node to the end of the list.
func (ll *LinkedList[V]) Append(value V) {
	newNode := NewNode(value)
	ll.nSize++

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		return
	}

	newNode.prev = ll.tail
	ll.tail.next = newNode
	ll.tail = ll.tail.next
}

// PushHead adds a new node to the head of the list.
func (ll *LinkedList[V]) PushHead(value V) {
	newNode := NewNode(value)
	ll.nSize++

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		return
	}

	newNode.next = ll.head
	ll.head.prev = newNode
	ll.head = newNode
}

// InsertAt adds a new node to the given position
func (ll *LinkedList[V]) InsertAt(pos int, value V) error {
	newNode := NewNode(value)

	if pos < 1 || uint(pos) > ll.nSize {
		return errors.New("invalid position")
	}

	currPos := 1
	currNode := ll.head

	if pos == 1 {
		ll.PushHead(value)
		return nil
	}

	for currPos < pos {
		currNode = currNode.next
		currPos++
	}

	newNode.prev = currNode.prev
	newNode.next = currNode
	currNode.prev.next = newNode
	currNode.prev = newNode

	ll.nSize++

	return nil
}

// Pop removes the last node from the list.
func (ll *LinkedList[V]) Pop() {
	if ll.head == nil {
		return
	}

	ll.nSize--

	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
		return
	}

	prev := ll.tail.prev
	prev.next = nil
	ll.tail = prev
}

// PopHead removes the first node from the list.
func (ll *LinkedList[V]) PopHead() {
	if ll.head == nil {
		return
	}

	ll.nSize--

	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
		return
	}

	next := ll.head.next
	next.prev = nil
	ll.head = next
}

// Find returns the node with the given value.
func (ll *LinkedList[V]) Find(value V) *Node[V] {
	curr := ll.head

	for curr != nil {
		if ll.cmp(curr.value, value) == 0 {
			return curr
		}

		curr = curr.next
	}

	return nil
}

// String returns a string representation of the list.
func (ll *LinkedList[V]) String() string {
	var s string

	curr := ll.head

	for curr != nil {
		s += fmt.Sprintf("[%v]", curr.value)
		if curr.next != nil {
			s += `-`
		}

		curr = curr.next
	}

	return s
}

// Head returns the head node of the list.
func (ll *LinkedList[V]) Head() *Node[V] {
	return ll.head
}

// Tail returns the tail node of the list.
func (ll *LinkedList[V]) Tail() *Node[V] {
	return ll.tail
}

// Size returns the size of the list.
func (ll *LinkedList[V]) Size() uint {
	return ll.nSize
}

// IsEmpty returns true if the list is empty.
func (ll *LinkedList[V]) IsEmpty() bool {
	return ll.head == nil
}
