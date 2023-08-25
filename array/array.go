package array

type Generic[T any] interface {
	Iterate(func(i int, elem T))
	At(i int) T
	RefAt(i int) *T
	Len() int
	Slice(slicer *Slicer) Slice[T]
}
