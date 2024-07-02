package lists

type ListIndexOutOfBoundsError struct {
	ind  uint
	from string
}

type ElementNotFoundError[T comparable] struct {
	el   T
	from string
}

type ListSizeTooSmallError string

type List[T comparable] interface {

	/*
		Adds the passed element at the passed index

		ind can range from 0 to l.Size()
	*/
	Add(uint, T) error

	// Deletes the element at the passed index
	Delete(uint) error

	// Adds the passed element to the end of the list and returns the new size of the list
	Push(T) uint

	// Deletes the last element of the list and returns the new size of the list
	Pop() (uint, error)

	// Returns the value of the element at the passed index
	Get(uint) (T, error)

	// Returns the index of the passed element
	IndexOf(T) (uint, error)

	// Returns the size of the list
	Size() uint
}
