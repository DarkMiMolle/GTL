package array

import (
	"fmt"
	"reflect"
)

// TODO: test

type ComparisonResult int

const (
	Before ComparisonResult = -1 + iota
	Same
	After
)

type ElemComparisonFunction[T any] func(left, right T) ComparisonResult

type Sorted[T any] struct {
	comparisonFunction ElemComparisonFunction[T]
	array              []T
}

func SortedWith[T any](comparisonFunction ElemComparisonFunction[T], elems ...T) (arr Sorted[T]) {
	arr.comparisonFunction = comparisonFunction
	arr.InsertMany(elems...)
	return
}
func (s *Sorted[T]) Insert(elem T) (idx int) {
	if len(s.array) == 0 {
		s.array = []T{elem}
		return 0
	}

	left := 0
	right := len(s.array) - 1
	arr := s.array
	for right-left > 1 {
		mid := (left + right + 1) / 2 // mid-distance between left and right (included) is (right+1 - left)/2;
		// but we want the index, so we need to add left: left + ((right+1 - left)/2);
		// which is = to: (left + right+1)/2
		if place := s.comparisonFunction(elem, arr[mid]); place <= Before { // elem is before mid
			right = mid
		} else if place >= After {
			left = mid
		} else { // Same
			left = mid
			break
		}
	}
	if s.comparisonFunction(elem, arr[left]) <= Before { // before left
		idx = left
		if idx == 0 {
			s.array = append([]T{elem}, arr...)
			return
		}
		s.array = append(arr[:left], append([]T{elem}, arr[left:]...)...)
		return idx
	}
	if s.comparisonFunction(elem, arr[right]) >= After { // after right
		idx = right + 1
		if right+1 >= len(arr) {
			s.array = append(arr, elem)
			return
		}
		s.array = append(arr[:right+1], append([]T{elem}, arr[:right+1]...)...)
		return idx
	}
	// between left and right
	idx = left + 1
	s.array = append(arr[:left+1], append([]T{elem}, arr[left+1:]...)...)
	return idx
}
func (s *Sorted[T]) InsertMany(elem ...T) (allIdx []int) {
	for _, elem := range elem {
		allIdx = append(allIdx, s.Insert(elem))
	}
	return allIdx
}
func (s Sorted[T]) String() string {
	return fmt.Sprint(s.array)
}
func (s *Sorted[T]) At(i int) T {
	return s.array[i]
}
func (s *Sorted[T]) RefAt(i int) *T {
	return &s.array[i]
}
func (s *Sorted[T]) Len() int {
	return len(s.array)
}
func (s *Sorted[T]) Iterate(forEach func(i int, elem T)) {
	for i, elem := range s.array {
		forEach(i, elem)
	}
}
func (s *Sorted[T]) Slice(slicer *Slicer) Slice[T] {
	return Slice[T]{
		array:  s,
		slicer: *slicer,
	}
}

func (s Sorted[T]) ComparisonFunction() ElemComparisonFunction[T] {
	return s.comparisonFunction
}
func (s Sorted[T]) UnderlyingArray() []T {
	arr := make([]T, len(s.array))
	copy(arr, s.array)
	return arr
}
func (s *Sorted[T]) First() T {
	return s.At(0)
}
func (s *Sorted[T]) Last() T {
	return s.At(s.Len() - 1)
}

func (s Sorted[T]) Find(elem T) (idx int, found bool) {
	arr := s.array

	left := 0
	right := s.Len() - 1
	idx = s.Len() / 2

	for right > left {
		if reflect.DeepEqual(elem, arr[idx]) {
			return idx, true
		}
		switch s.comparisonFunction(elem, arr[idx]) {
		case Before:
			right = idx - 1
		case After:
			left = idx + 1
		case Same:
			return idx, true
		}
		idx = (left + right + 1) / 2
		if idx >= s.Len() {
			idx = s.Len() - 1
			break
		}
		if idx < 0 {
			idx = 0
			break
		}
	}
	return idx, reflect.DeepEqual(elem, arr[idx])
}

func (s Sorted[T]) SplitAtIdx(i int) (left, right Sorted[T]) {
	left = s
	left.array = make([]T, len(s.array[:i]))
	copy(left.array, s.array)

	right = s
	right.array = make([]T, len(s.array[i:]))
	copy(right.array, s.array[i:])

	return
}
func (s Sorted[T]) SplitAtElem(elem T) (left, right Sorted[T]) {
	if s.array == nil {
		return s, s
	}

	if s.comparisonFunction(elem, s.Last()) >= After {
		return s.SplitAtIdx(s.Len())
	}

	idx, found := s.Find(elem)
	if found {
		return s.SplitAtIdx(idx + 1)
	}
	return s.SplitAtIdx(idx)
}
func (s Sorted[T]) SplitWhen(cond func(elem T) bool) (left, right Sorted[T]) {
	left.comparisonFunction = s.comparisonFunction
	right.comparisonFunction = s.comparisonFunction

	for idx, elem := range s.array {
		if cond(elem) {
			return s.SplitAtIdx(idx)
		}
	}
	left.array = s.UnderlyingArray()
	return
}
