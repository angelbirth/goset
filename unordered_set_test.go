package set

import (
	"reflect"
	"testing"
)

func Test_unorderedSet_Add(t *testing.T) {
	type testCase[T comparable] struct {
		name    string
		setFunc func() unorderedSet[T]
		arg     T
		want    bool
	}
	tests := []testCase[string]{
		{
			name: "add to empty set",
			setFunc: func() unorderedSet[string] {
				return unorderedSet[string]{}
			},
			arg:  "a",
			want: true,
		},
		{
			name: "add duplicate element",
			setFunc: func() unorderedSet[string] {
				s := unorderedSet[string]{}
				s.Add("a")
				return s
			},
			arg:  "a",
			want: false,
		},
		{
			name: "add distinct element",
			setFunc: func() unorderedSet[string] {
				s := unorderedSet[string]{}
				s.Add("a")
				return s
			},
			arg:  "s",
			want: true,
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

func Test_unorderedSet_Clear(t *testing.T) {
	type testCase[T comparable] struct {
		name    string
		setFunc func() unorderedSet[T]
	}
	tests := []testCase[string]{
		{
			name: "clear non-empty set",
			setFunc: func() unorderedSet[string] {
				set := unorderedSet[string]{}
				set.Add("a")
				return set
			},
		},
		{
			name: "clear empty set",
			setFunc: func() unorderedSet[string] {
				return unorderedSet[string]{}
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

func Test_unorderedSet_Clone(t *testing.T) {
	type testCase[T comparable] struct {
		name     string
		setFunc  func() unorderedSet[T]
		wantFunc func() Set[T]
	}
	tests := []testCase[string]{
		// TODO: Add test cases.
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

func Test_unorderedSet_Contains(t *testing.T) {
	type args[T comparable] struct {
		val T
	}
	type testCase[T comparable] struct {
		name string
		s    unorderedSet[T]
		args args[T]
		want bool
	}
	tests := []testCase[string]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.val); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unorderedSet_Remove(t *testing.T) {
	type args[T comparable] struct {
		val T
	}
	type testCase[T comparable] struct {
		name  string
		s     unorderedSet[T]
		args  args[T]
		want  T
		want1 bool
	}
	tests := []testCase[string]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Remove(tt.args.val)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Remove() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_unorderedSet_Size(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    unorderedSet[T]
		want int
	}
	tests := []testCase[string]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unorderedSet_Slice(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    unorderedSet[T]
		want []T
	}
	tests := []testCase[string]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
