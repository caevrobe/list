package list

import "fmt"

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
}

type Node[T any] struct {
	Value T
	prev  *Node[T]
	next  *Node[T]
}

/* List functions *************************************************************************/

// returns pointer to new typed List
func New[T any]() *List[T] {
	return &List[T]{}
}

// returns pointer to list head
func (l *List[T]) Head() *Node[T] {
	return l.head
}

// inserts new Node with given Value. returns pointer to new Node
func (l *List[T]) Insert(Value T) *Node[T] {
	obj := &Node[T]{Value, nil, nil}

	if l.head == nil {
		l.head = obj
	} else {
		l.tail.next = obj
	}

	obj.prev = l.tail
	l.tail = obj
	return obj
}

// removes given Node from List
func (l *List[T]) Remove(obj *Node[T]) {
	if obj == l.head {
		l.head = obj.next
	}
	if obj == l.tail {
		l.tail = obj.prev
	}

	if obj.prev != nil {
		obj.prev.next = obj.next
	}
	if obj.next != nil {
		obj.next.prev = obj.prev
	}

	obj.prev = nil
	obj.next = nil
}

// gross
func (l *List[T]) String() string {
	contents := "["
	for item := l.head; item != nil; item = item.next {
		contents += fmt.Sprint(item.Value)
		if item.next != nil {
			contents += ", "
		}
	}
	contents += "]"

	h_val, t_val := "nil", "nil"
	if l.head != nil {
		h_val = fmt.Sprint(l.head.Value)
	}
	if l.tail != nil {
		t_val = fmt.Sprint(l.tail.Value)
	}

	return "H: " + h_val + " | T: " + t_val + " | " + contents
}

/* Node functions *************************************************************************/

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}
