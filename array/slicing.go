package array

import (
	"fmt"
	"strings"

	. "github.com/DarkMiMolle/GTL/utils"
)

// TODO: test

type Slicer struct {
	from Opt[int]
	to   Opt[int]
	step Opt[int]
}

func (s *Slicer) From(i int) *Slicer {
	s.from.Set(i)
	return s
}
func (s *Slicer) To(i int) *Slicer {
	s.to.Set(i)
	return s
}
func (s *Slicer) WithStepOf(i int) *Slicer {
	s.step.Set(i)
	return s
}
func (s Slicer) ConvertIndex(inputIdx int, arrayLen int) (arrayIdx int, err error) {
	from := s.from.ValueOr(0)
	step := s.step.ValueOr(1)
	if !s.from.HasValue() && step < 0 {
		from = -1
	}
	idx := from + s.step.ValueOr(1)*inputIdx
	if idx < 0 {
		idx = arrayLen + idx // + idx <=> - unsigned value of idx (like a + (-4) <=> a - 4)
	}
	if idx < 0 || idx >= arrayLen {
		return 0, fmt.Errorf("index out of range")
	}
	return idx, nil
}
func (s Slicer) ConvertLen(arrayLen int) (l int) {
	defer func() {
		if l < 0 {
			l = 0
		}
	}()
	step := s.step.ValueOr(1)
	if step < 0 {
		to := s.to
		s.to.Set(s.from.ValueOr(0) + 1)
		s.from.Set(to.ValueOr(-1) + 1)
		step = -step
	}
	return (s.to.ValueOr(arrayLen) - s.from.ValueOr(0)) / step
}

type Slice[T any] struct {
	array  Generic[T]
	slicer Slicer
}

func (s Slice[T]) Iterate(forEach func(i int, elem T)) {
	for i := 0; i < s.Len(); i++ {
		forEach(i, s.At(i))
	}
}
func (s Slice[T]) At(i int) T {
	idx, err := s.slicer.ConvertIndex(i, s.array.Len())
	if err != nil {
		panic(err)
	}
	return s.array.At(idx)
}
func (s Slice[T]) Len() int {
	return s.slicer.ConvertLen(s.array.Len())
}
func (s Slice[T]) RefAt(i int) *T {
	idx, err := s.slicer.ConvertIndex(i, s.Len())
	if err != nil {
		panic(err)
	}
	return s.array.RefAt(idx)
}
func (s Slice[T]) Slice(slicer *Slicer) Slice[T] {
	return Slice[T]{
		array:  s,
		slicer: *slicer,
	}
}
func (s Slice[T]) String() string {
	str := "["
	s.Iterate(func(_ int, elem T) {
		str += fmt.Sprint(elem) + ", "
	})
	str = strings.TrimSuffix(str, ", ")
	return str + "]"
}
