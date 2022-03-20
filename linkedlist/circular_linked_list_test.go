package linkedlist_test

import (
	"testing"

	"github.com/HotPotatoC/sture/linkedlist"
)

func TestCircularLinkedList_Append(t *testing.T) {
	ll := linkedlist.NewCircularLinkedList[int]()
	exp := "[1]-[2]-[3]-[4]-[5]"

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)
	ll.Append(5)

	if ll.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, ll.String())
	}
}

func TestCircularLinkedList_PushHead(t *testing.T) {
	ll := linkedlist.NewCircularLinkedList[int]()
	exp := "[5]-[4]-[3]-[2]-[1]"

	ll.PushHead(1)
	ll.PushHead(2)
	ll.PushHead(3)
	ll.PushHead(4)
	ll.PushHead(5)

	if ll.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, ll.String())
	}
}

func TestCircularLinkedList_Pop(t *testing.T) {
	ll := linkedlist.NewCircularLinkedList[int]()
	exp := "[1]-[2]-[5]"

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Pop()
	ll.Append(4)
	ll.Pop()
	ll.Append(5)

	if ll.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, ll.String())
	}

	exp = "[1]"

	ll.Pop()
	ll.Pop()

	if ll.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, ll.String())
	}

	exp = ""

	ll.Pop()

	if ll.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, ll.String())
	}
}

func TestCircularLinkedList_PopHead(t *testing.T) {
	ll := linkedlist.NewCircularLinkedList[int]()
	exp := "[5]-[2]-[1]"

	ll.PushHead(1)
	ll.PushHead(2)
	ll.PushHead(3)
	ll.PopHead()
	ll.PushHead(4)
	ll.PopHead()
	ll.PushHead(5)

	if ll.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, ll.String())
	}

	exp = "[1]"

	ll.PopHead()
	ll.PopHead()

	if ll.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, ll.String())
	}

	exp = ""

	ll.PopHead()

	if ll.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, ll.String())
	}
}

func TestCircularLinkedList_Find(t *testing.T) {
	ll := linkedlist.NewCircularLinkedList[int]()

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	node := ll.Find(2)

	if node.Next().Value() != 3 || node.Prev().Value() != 1 {
		t.Errorf("\nExpected: next=3 prev=1\nGot: next=%d prev=%d", node.Next().Value(), node.Prev().Value())
	}
}

func TestCircularLinkedList_IsCircular(t *testing.T) {
	ll := linkedlist.NewCircularLinkedList[int]()

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	if ll.Head().Prev() != ll.Tail() {
		t.Errorf("\nExpected: prev=tail\nGot: prev=%d", ll.Head().Prev().Value())
	}

	if ll.Tail().Next() != ll.Head() {
		t.Errorf("\nExpected: next=head\nGot: next=%d", ll.Tail().Next().Value())
	}

	if !ll.IsCircular() {
		t.Errorf("\nExpected: circular")
	}
}
