package linkedlist

import (
	"errors"
	"fmt"
)

// LinkedList is a linked list.
type LinkedList[T any] struct {
	head  *Node[T]
	tail  *Node[T]
	nSize uint
	cmp   func(T, T) int
}

// NewLinkedList returns a new linked list.
func NewLinkedList[T any](cmp func(T, T) int) *LinkedList[T] {
	return &LinkedList[T]{cmp: cmp}
}

// Append adds a new node to the end of the list.
func (ll *LinkedList[T]) Append(value T) {
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
func (ll *LinkedList[T]) PushHead(value T) {
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
func (ll *LinkedList[T]) InsertAt(pos int, value T) error {
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
func (ll *LinkedList[T]) Pop() {
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
func (ll *LinkedList[T]) PopHead() {
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
func (ll *LinkedList[T]) Find(value T) *Node[T] {
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
func (ll *LinkedList[T]) String() string {
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
func (ll *LinkedList[T]) Head() *Node[T] {
	return ll.head
}

// Tail returns the tail node of the list.
func (ll *LinkedList[T]) Tail() *Node[T] {
	return ll.tail
}

// Size returns the size of the list.
func (ll *LinkedList[T]) Size() uint {
	return ll.nSize
}

// IsEmpty returns true if the list is empty.
func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.head == nil
}
