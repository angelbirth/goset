package set

type unorderedSet[T comparable] map[T]struct{}

func (s unorderedSet[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s unorderedSet[T]) Clone() Set[T] {
	newSet := make(unorderedSet[T], len(s))
	for val := range s {
		newSet[val] = struct{}{}
	}
	return newSet
}

func (s unorderedSet[T]) Add(val T) bool {
	if _, exists := s[val]; exists {
		return false
	}
	s[val] = struct{}{}
	return true
}

func (s unorderedSet[T]) Remove(val T) (T, bool) {
	_, exists := s[val]
	if exists {
		delete(s, val)
		return val, true
	}
	var v T
	return v, false
}

func (s unorderedSet[T]) Size() int {
	return len(s)
}

func (s unorderedSet[T]) Slice() []T {
	res := make([]T, len(s))
	for k := range s {
		res = append(res, k)
	}
	return res
}

func (s unorderedSet[T]) Contains(val T) bool {
	_, exists := s[val]
	return exists
}
