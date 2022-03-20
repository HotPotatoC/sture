package linkedlist

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

// CircularLinkedList is a circular linked list.
type CircularLinkedList[V constraints.Ordered] struct {
	head *Node[V]
	tail *Node[V]
	nSize uint
}

// NewCircularLinkedList returns a new circular linked list.
func NewCircularLinkedList[V constraints.Ordered]() *CircularLinkedList[V] {
	return &CircularLinkedList[V]{}
}

// Append adds a new node to the end of the list.
func (ll *CircularLinkedList[V]) Append(value V) {
	newNode := NewNode(value)
	ll.nSize++

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		return
	}

	newNode.prev = ll.tail
	newNode.next = ll.head
	ll.tail.next = newNode
	ll.tail = ll.tail.next
	ll.head.prev = ll.tail
}

// PushHead adds a new node to the head of the list.
func (ll *CircularLinkedList[V]) PushHead(value V) {
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
	ll.tail.next = ll.head
}

// InsertAt adds a new node to the given position
func (ll *CircularLinkedList[V]) InsertAt(pos int, value V) error {
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

// PushMid adds a new node to the middle of the list.
func (ll *CircularLinkedList[V]) PushMid(value V) {
	newNode := NewNode(value)
	ll.nSize++

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
		return
	}

	if ll.head.value > value {
		ll.PushHead(value)
		return
	}

	if ll.tail.value < value {
		ll.Append(value)
		return
	}

	curr := ll.head

	for curr.value < value {
		curr = curr.next
	}

	newNode.prev = curr.prev
	newNode.next = curr
	curr.prev.next = newNode
	curr.prev = newNode
}

// Pop removes the last node from the list.
func (ll *CircularLinkedList[V]) Pop() {
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
	prev.next = ll.head
	ll.tail = prev
	ll.head.prev = ll.tail
}

// PopHead removes the first node from the list.
func (ll *CircularLinkedList[V]) PopHead() {
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
	next.prev = ll.tail
	ll.head = next
	ll.tail.next = ll.head
}

// Find returns the node with the given value.
func (ll *CircularLinkedList[V]) Find(value V) *Node[V] {
	curr := ll.head

	for curr != nil {
		if curr.value == value {
			return curr
		}

		curr = curr.next
	}

	return nil
}

// String returns a string representation of the list.
func (ll *CircularLinkedList[V]) String() string {
	var s string

	curr := ll.head

	for curr != nil {
		s += fmt.Sprintf("[%v]", curr.value)
		if curr.next != ll.head {
			s += `-`
		}

		if curr.next == ll.head {
			return s
		}

		curr = curr.next
	}

	return s
}

// Head returns the head node of the list.
func (ll *CircularLinkedList[V]) Head() *Node[V] {
	return ll.head
}

// Tail returns the tail node of the list.
func (ll *CircularLinkedList[V]) Tail() *Node[V] {
	return ll.tail
}

// Size returns the size of the list.
func (ll *CircularLinkedList[V]) Size() uint {
	return ll.nSize
}

// IsCircular returns true if the list is circular.
func (ll *CircularLinkedList[V]) IsCircular() bool {
	return ll.head != nil && ll.head == ll.tail.next
}

// IsEmpty returns true if the list is empty.
func (ll *CircularLinkedList[V]) IsEmpty() bool {
	return ll.head == nil
}
