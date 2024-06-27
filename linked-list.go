package main

import (
	"fmt"
)

type Node[T any] struct {
	next *Node[T]
	val  T
}

type List[T any] struct {
	head *Node[T]
}

type ListIndexOutOfBoundsError int

func (e ListIndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("index %v is invalid: index is out of bounds", int(e))
}

func (l List[T]) String() (s string) {
	s = "["
	for i := range l.Size() {
		el, err := l.Get(i)
		if err != nil {
			return
		}

		s += fmt.Sprintf("%v", el)

		if i < l.Size()-1 {
			s += ", "
		}
	}

	s += "]"
	return
}

func (l *List[T]) Size() (i uint) {
	iter := l.head
	if iter == nil {
		return uint(0)
	}

	i = uint(1)
	for ; ; i++ {
		iter = iter.next

		if iter.next == nil {
			i++
			break
		}
	}

	return
}

func (l *List[T]) Get(ind uint) (e T, err error) {
	if ind >= l.Size() {
		err = ListIndexOutOfBoundsError(ind)
		return
	}

	iter := l.head
	for i := uint(0); ; i, iter = i+1, iter.next {
		if i == ind {
			e = iter.val
			break
		}

		if iter.next == nil {
			break
		}
	}

	return
}

func (l *List[T]) Add(e T) {
	iter := l.head
	for i := 0; ; i, iter = i+1, iter.next {
		if iter.next == nil {
			iter.next = &Node[T]{nil, e}
			break
		}
	}
}

func main() {
	node3 := Node[int]{nil, 3}
	node2 := Node[int]{&node3, 2}
	node1 := Node[int]{&node2, 1}

	list := List[int]{&node1}
	list.Add(10)
	list.Add(15)
	fmt.Println(list)

	list2 := List[int]{}
	fmt.Println(list2)
}
