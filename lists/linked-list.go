package lists

import (
	"fmt"
)

type node[T any] struct {
	next *node[T]
	val  T
}

type linkedList[T any] struct {
	head *node[T]
}

// Initialize a linked list with an array
func LinkedList[T any](a []T) *linkedList[T] {
	l := &linkedList[T]{}
	for i := 0; i < len(a); i++ {
		l.Add(a[i])
	}

	return l
}

func (l *linkedList[T]) Add(e T) {
	new := &node[T]{nil, e}

	if l.head == nil {
		l.head = new
		return
	}

	iter := l.head
	for iter.next != nil {
		iter = iter.next
	}

	iter.next = new
}

// Deletes the element at the passed index
func (l *linkedList[T]) Delete(ind uint) (err error) {
	var el, prev *node[T]
	for i := range l.Size() {
		el, err = l.GetNode(i)

		switch {
		case err != nil:
			return

		case i == ind && i == 0:
			l.head = el.next
			return

		case i == ind:
			prev, err = l.GetNode(i - 1)

			if err != nil {
				return
			}

			prev.next = el.next
			return
		}
	}

	err = ElementNotFoundError(ind)
	return
}

func (l *linkedList[T]) Get(ind uint) (e T, err error) {
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

func (l *linkedList[T]) GetNode(ind uint) (e *node[T], err error) {
	if ind >= l.Size() {
		err = ListIndexOutOfBoundsError(ind)
		return
	}

	iter := l.head
	for i := uint(0); ; i, iter = i+1, iter.next {
		if i == ind {
			e = iter
			break
		}

		if iter.next == nil {
			break
		}
	}

	return
}

func (l *linkedList[T]) Size() (i uint) {
	iter := l.head

	if iter == nil {
		return
	}

	for i = 1; iter.next != nil; i++ {
		iter = iter.next
	}

	return
}

func (e ListIndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("index %v is invalid: index is out of bounds", uint(e))
}

func (e ElementNotFoundError) Error() string {
	return fmt.Sprintf("could not delete element at index %v: element not found", uint(e))
}

func (l linkedList[T]) String() (s string) {
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
