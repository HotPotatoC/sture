package linkedlist_test

import (
	"testing"

	"github.com/HotPotatoC/sture"
	"github.com/HotPotatoC/sture/linkedlist"
)

func TestLinkedList_Append(t *testing.T) {
	ll := linkedlist.NewLinkedList(sture.Compare[int])
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

func TestLinkedList_PushHead(t *testing.T) {
	ll := linkedlist.NewLinkedList(sture.Compare[int])
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

func TestLinkedList_InsertAt(t *testing.T) {
	ll := linkedlist.NewLinkedList(sture.Compare[int])

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	ll.InsertAt(2, 4)

	exp := "[1]-[4]-[2]-[3]"

	if ll.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, ll.String())
	}

	ll.InsertAt(1, 5)

	exp = "[5]-[1]-[4]-[2]-[3]"

	if ll.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, ll.String())
	}

	ll.InsertAt(5, 6)

	exp = "[5]-[1]-[4]-[2]-[6]-[3]"

	if ll.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, ll.String())
	}

	err := ll.InsertAt(7, 7)
	if err == nil {
		t.Errorf("\nExpected: error\nGot: nil")
	}
}

func TestLinkedList_Pop(t *testing.T) {
	ll := linkedlist.NewLinkedList(sture.Compare[int])
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

func TestLinkedList_PopHead(t *testing.T) {
	ll := linkedlist.NewLinkedList(sture.Compare[int])
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

func TestLinkedList_Find(t *testing.T) {
	ll := linkedlist.NewLinkedList(sture.Compare[int])

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	node := ll.Find(2)

	if node.Next().Value() != 3 || node.Prev().Value() != 1 {
		t.Errorf("\nExpected: next=3 prev=1\nGot: next=%d prev=%d",
			node.Next().Value(), node.Prev().Value())
	}
}
