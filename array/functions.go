package array

// MapRef allows to generate a new slice from reference's value of an existing one. If you don't need reference use Map
func MapRef[T1, T2 any](arr []T1, action func(elem *T1) T2) []T2 {
	newArr := make([]T2, len(arr))
	for idx := range arr {
		newArr[idx] = action(&arr[idx])
	}
	return newArr
}

// Map allows to generate a new slice from value of an existing one. If you need to get references of the existing slice use MapRef
func Map[T1, T2 any](arr []T1, action func(elem T1) T2) []T2 {
	newArr := make([]T2, len(arr))
	for idx, value := range arr {
		newArr[idx] = action(value)
	}
	return newArr
}

// Reversed reverse the given slice
func Reversed[T any](arr []T) []T {
	rev := make([]T, len(arr))
	for i := 0; i < len(arr)/2; i++ {
		rev[i] = arr[len(arr)-1-i]
		rev[len(arr)-1-i] = arr[i]
	}
	return rev
}

// Break is a way to represent the break statement in a for loop; true mean breaking the loop
// this is to use as a function return type.
type Break bool

// ForEachReverse apply action like a for-range loop on a slice, in reversed order
func ForEachReverse[T any](arr []T, action func(elem T) (breaking Break)) {
	for idx := range arr {
		if action(arr[len(arr)-1-idx]) {
			return
		}
	}
}

// ForEach apply action like a for-range loop on a slice
func ForEach[T any](arr []T, action func(elem T) (breaking Break)) {
	for _, elem := range arr {
		if action(elem) {
			return
		}
	}
}

// Filter will generate a new slice containing only element that match the filter function (a match mean the function returns true)
func Filter[T any](arr []T, filter func(elem T) bool) []T {
	filtered := make([]T, 0, len(arr))
	for _, val := range arr {
		if filter(val) {
			filtered = append(filtered, val)
		}
	}
	return filtered
}

func FindElemRef[T comparable](arr []T, elem T) (*T, bool) {
	for i := range arr {
		ref := &arr[i]
		if *ref == elem {
			return ref, true
		}
	}
	return nil, false
}

func FindElem[T comparable](arr []T, elem T) (int, bool) {
	for i, val := range arr {
		if val == elem {
			return i, true
		}
	}
	return 0, false
}

func FindMatchRef[T any](arr []T, match func(T) bool) (*T, bool) {
	for i, val := range arr {
		if match(val) {
			return &arr[i], true
		}
	}
	return nil, false
}

func FindMatch[T any](arr []T, match func(T) bool) (int, bool) {
	for i, val := range arr {
		if match(val) {
			return i, true
		}
	}
	return 0, false
}

func ReduceAs[T1, T2 any](arr []T1, reduce func(reduced T2, elem T1) T2) T2 {
	var result T2
	for _, elem := range arr {
		result = reduce(result, elem)
	}
	return result
}

func Reduce[T any](arr []T, reduce func(reduced, elem T) T) T {
	return ReduceAs(arr, reduce)
}
