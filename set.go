package set

type Set[T comparable] interface {
	Add(val T) bool
	Remove(val T) (T, bool)
	Clear()
	Size() int
	Slice() []T
	Contains(val T) bool
	Clone() Set[T]
}
