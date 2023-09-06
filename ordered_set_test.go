package set

import (
	"reflect"
	"testing"
)

func Test_orderedSet_Add(t *testing.T) {
	type testCase[T comparable] struct {
		name    string
		setFunc func() *orderedSet[T]
		arg     T
		want    bool
	}
	tests := []testCase[string]{
		{
			name: "add to empty set",
			setFunc: func() *orderedSet[string] {
				return newOrderedSet[string]()
			},
			arg:  "arg1",
			want: true,
		},
		{
			name: "add existing element",
			setFunc: func() *orderedSet[string] {
				set := newOrderedSet[string]()
				set.Add("arg1")
				return set
			},
			arg:  "arg1",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.setFunc().Add(tt.arg); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderedSet_Clear(t *testing.T) {
	type testCase[T comparable] struct {
		name    string
		setFunc func() *orderedSet[T]
	}

	tests := []testCase[string]{
		{
			name: "clear non-empty set",
			setFunc: func() *orderedSet[string] {
				set := newOrderedSet[string]()
				set.Add("a")
				return set
			},
		},
		{
			name: "clear empty set",
			setFunc: func() *orderedSet[string] {
				return newOrderedSet[string]()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := tt.setFunc()
			set.Clear()
			if set.Size() > 0 {
				t.Errorf("Clear(): got Size()=%d, want Size()=%d", 0, set.Size())
			}
		})
	}
}

func Test_orderedSet_Clone(t *testing.T) {
	type testCase[T comparable] struct {
		name     string
		setFunc  func() *orderedSet[T]
		wantFunc func() Set[T]
	}

	tests := []testCase[string]{
		{
			name: "clone",
			setFunc: func() *orderedSet[string] {
				set := newOrderedSet[string]()
				set.Add("element1")
				set.Add("element2")
				return set
			},
			wantFunc: func() Set[string] {
				set := newOrderedSet[string]()
				set.Add("element1")
				set.Add("element2")
				return set
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := tt.setFunc()
			got, want := set.Clone(), tt.wantFunc()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Clone() = %v, want %v", got, want)
			}
		})
	}
}

func Test_orderedSet_Contains(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		o    *orderedSet[T]
		arg  T
		want bool
	}
	set := newOrderedSet[string]()
	set.Add("arg1")
	tests := []testCase[string]{
		{
			name: "exists",
			o:    set,
			arg:  "arg1",
			want: true,
		},
		{
			name: "not exists",
			o:    set,
			arg:  "arg2",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Contains(tt.arg); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderedSet_Remove(t *testing.T) {
	type testCase[T comparable] struct {
		name  string
		o     *orderedSet[T]
		arg   T
		want  T
		want1 bool
	}
	set := newOrderedSet[string]()
	set.Add("a")
	set.Add("b")
	tests := []testCase[string]{
		{
			name:  "remove existing",
			o:     set,
			arg:   "a",
			want:  "a",
			want1: true,
		},
		{
			name:  "remove non-existing",
			o:     set,
			arg:   "c",
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.o.Remove(tt.arg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Remove() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_orderedSet_Size(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		o    *orderedSet[T]
		want int
	}
	set := newOrderedSet[string]()
	set.Add("a")
	tests := []testCase[string]{
		{
			name: "empty set",
			o:    newOrderedSet[string](),
			want: 0,
		},
		{
			name: "single element",
			o:    set,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderedSet_Slice(t *testing.T) {
	type testCase[T comparable] struct {
		name     string
		set      *orderedSet[T]
		want     []T
		scenario func(set *orderedSet[string])
	}
	tests := []testCase[string]{
		{
			name:     "empty set",
			set:      newOrderedSet[string](),
			want:     nil,
			scenario: func(set *orderedSet[string]) {},
		},
		{
			name: "non-empty set",
			set:  newOrderedSet[string](),
			want: []string{"a", "s"},
			scenario: func(set *orderedSet[string]) {
				set.Add("a")
				set.Add("s")
			},
		},
		{
			name: "add 2 of the same element",
			set:  newOrderedSet[string](),
			want: []string{"a"},
			scenario: func(set *orderedSet[string]) {
				set.Add("a")
				set.Add("a")
			},
		},
		{
			name: "add 2 elements, remove 1, then add another",
			set:  newOrderedSet[string](),
			want: []string{"s", "d"},
			scenario: func(set *orderedSet[string]) {
				set.Add("a")
				set.Add("s")
				set.Remove("a")
				set.Add("d")
			},
		},
		{
			name: "remove element from single element set",
			set:  newOrderedSet[string](),
			want: []string{},
			scenario: func(set *orderedSet[string]) {
				set.Add("a")
				set.Remove("a")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.scenario(tt.set)
			if got := tt.set.Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
