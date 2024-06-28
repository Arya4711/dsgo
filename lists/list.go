package lists

type ListIndexOutOfBoundsError uint
type ElementNotFoundError uint

type List[T any] interface {
	Add(T)
	Delete(uint) error
	// Push(T)
	// Pop(T)
	Get(uint) (T, error)
	// IndexOf(T) uint
	Size() uint
}
