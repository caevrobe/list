package list

import "fmt"

type list[T any] struct {
	head *node[T]
	tail *node[T]
}

type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

func New[T any]() *list[T] {
	return &list[T]{}
}

func (l *list[T]) Insert(value T) *node[T] {
	obj := &node[T]{value, nil, nil}

	if l.head == nil {
		l.head = obj
	} else {
		l.tail.next = obj
	}

	obj.prev = l.tail
	l.tail = obj
	return obj
}

func (l *list[T]) Remove(obj *node[T]) {
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
}

// gross
func (l *list[T]) String() string {
	contents := "["
	for item := l.head; item != nil; item = item.next {
		contents += fmt.Sprint(item.value)
		if item.next != nil {
			contents += ", "
		}
	}
	contents += "]"

	h_val, t_val := "nil", "nil"
	if l.head != nil {
		h_val = fmt.Sprint(l.head.value)
	}
	if l.tail != nil {
		t_val = fmt.Sprint(l.tail.value)
	}

	return "H: " + h_val + " | T: " + t_val + " | " + contents
}
