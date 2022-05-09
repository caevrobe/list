package list

import "fmt"

type List[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

type Node[T any] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}

func New[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Insert(Value T) *Node[T] {
	obj := &Node[T]{Value, nil, nil}

	if l.Head == nil {
		l.Head = obj
	} else {
		l.Tail.Next = obj
	}

	obj.Prev = l.Tail
	l.Tail = obj
	return obj
}

func (l *List[T]) Remove(obj *Node[T]) {
	if obj == l.Head {
		l.Head = obj.Next
	}
	if obj == l.Tail {
		l.Tail = obj.Prev
	}

	if obj.Prev != nil {
		obj.Prev.Next = obj.Next
	}
	if obj.Next != nil {
		obj.Next.Prev = obj.Prev
	}
}

// gross
func (l *List[T]) String() string {
	contents := "["
	for item := l.Head; item != nil; item = item.Next {
		contents += fmt.Sprint(item.Value)
		if item.Next != nil {
			contents += ", "
		}
	}
	contents += "]"

	h_val, t_val := "nil", "nil"
	if l.Head != nil {
		h_val = fmt.Sprint(l.Head.Value)
	}
	if l.Tail != nil {
		t_val = fmt.Sprint(l.Tail.Value)
	}

	return "H: " + h_val + " | T: " + t_val + " | " + contents
}
