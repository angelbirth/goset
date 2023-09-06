package set

import "slices"

type orderedSet[T comparable] struct {
	set  unorderedSet[T]
	list []T
}

func newOrderedSet[T comparable]() *orderedSet[T] {
	return &orderedSet[T]{
		set: make(unorderedSet[T]),
	}
}

func (o *orderedSet[T]) Add(val T) bool {
	if o.set.Add(val) {
		o.list = append(o.list, val)
		return true
	}
	return false
}

func (o *orderedSet[T]) Remove(val T) (T, bool) {
	_, exists := o.set[val]
	if exists {
		delete(o.set, val)
		i := slices.Index(o.list, val)
		if i != -1 {
			if i < len(o.list)-1 {
				o.list = append(o.list[:i], o.list[i+1:]...)
			} else {
				o.list = o.list[:i]
			}
		} else {
			panic("should not happen")
		}
		return val, true
	}
	var ret T
	return ret, false
}

func (o *orderedSet[T]) Clear() {
	o.list = nil
	for key := range o.set {
		delete(o.set, key)
	}
}

func (o *orderedSet[T]) Size() int {
	return len(o.set)
}

func (o *orderedSet[T]) Slice() []T {
	return o.list
}

func (o *orderedSet[T]) Contains(val T) bool {
	_, exists := o.set[val]
	return exists
}

func (o *orderedSet[T]) Clone() Set[T] {
	newSet := &orderedSet[T]{
		set:  make(unorderedSet[T], len(o.set)),
		list: make([]T, len(o.list)),
	}
	for key := range o.set {
		newSet.set[key] = struct{}{}
	}
	copy(newSet.list, o.list)
	return newSet
}
