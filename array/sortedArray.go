package array

import (
	"fmt"
)

// TODO: test

type ComparisonResult int

const (
	Before ComparisonResult = -1 + iota
	Same
	After
)

type ElemComparisonFunction[T any] func(left, right T) ComparisonResult

type IndexComparisonFunction[T any] func(arr []T, i, j int) ComparisonResult

type ComparisonFunction[T any] interface {
	ElemComparisonFunction[T] | IndexComparisonFunction[T]
}

type SortedWithFunction[T any, F ComparisonFunction[T]] struct {
	comparisonFunction F
	array              []T
}

func SortedWith[T any, F ComparisonFunction[T]](comparisonFunction F, elems ...T) (arr SortedWithFunction[T, F]) {
	defer arr.InsertMultiple(elems...)
	return SortedWithFunction[T, F]{
		comparisonFunction: comparisonFunction,
	}
}
func (s *SortedWithFunction[T, F]) callComparison(i, j int) (result ComparisonResult) {
	var comparisonFunction any = s.comparisonFunction
	switch comparisonFunction := comparisonFunction.(type) {
	case ElemComparisonFunction[T]:
		return comparisonFunction(s.array[i], s.array[j])
	case IndexComparisonFunction[T]:
		return comparisonFunction(s.array, i, j)
	}
	// s.comparisonFunction MUST be one of these 2 function type (by design)
	panic("unreachable")
}
func (s *SortedWithFunction[T, F]) Insert(elem T) (idx int) {
	if len(s.array) == 0 {
		s.array = []T{elem}
		return 0
	}
	left := 0
	right := len(s.array) - 1
	idx = len(s.array)
	arr := s.array                  // the array without the elem
	s.array = append(s.array, elem) // elem is at idx
	for right-left > 1 {
		mid := (left + right + 1) / 2 // mid-distance between left and right (included) is (right+1 - left)/2;
		// but we want the index, so we need to add left: left + ((right+1 - left)/2);
		// which is = to: (left + right+1)/2
		if place := s.callComparison(idx, mid); place <= Before { // elem is before mid
			right = mid
		} else if place >= After {
			left = mid
		} else { // Same
			left = mid
			break
		}
	}
	if s.callComparison(idx, left) <= Before { // before left
		idx = left
		if idx == 0 {
			s.array = append([]T{elem}, arr...)
			return
		}
		s.array = append(arr[:left], append([]T{elem}, arr[left:]...)...)
		return idx
	}
	if s.callComparison(idx, right) >= After { // after right
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
func (s *SortedWithFunction[T, F]) InsertMultiple(elem ...T) (allIdx []int) {
	for _, elem := range elem {
		allIdx = append(allIdx, s.Insert(elem))
	}
	return allIdx
}
func (s *SortedWithFunction[T, F]) String() string {
	return fmt.Sprint(s.array)
}
func (s *SortedWithFunction[T, F]) At(i int) T {
	return s.array[i]
}
func (s *SortedWithFunction[T, F]) RefAt(i int) *T {
	return &s.array[i]
}
func (s *SortedWithFunction[T, F]) Len() int {
	return len(s.array)
}
func (s *SortedWithFunction[T, F]) Iterate(forEach func(i int, elem T)) {
	for i, elem := range s.array {
		forEach(i, elem)
	}
}
func (s *SortedWithFunction[T, F]) Slice(slicer *Slicer) Slice[T] {
	return Slice[T]{
		array:  s,
		slicer: *slicer,
	}
}
