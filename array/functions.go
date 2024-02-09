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

// FindElemRef will return a reference to the first element from 'arr' equal to 'elem' (using '==') along with a boolean that indicate if the element is found
// if the element is not found the boolean is false and the pointer is nil
func FindElemRef[T comparable](arr []T, elem T) (*T, bool) {
	for i := range arr {
		ref := &arr[i]
		if *ref == elem {
			return ref, true
		}
	}
	return nil, false
}

// FindElem will return the index to the first element from 'arr' where 'arr[i] == elem' (with 'i' the index)
// it also returns a boolean set to true if the element is found and false otherwise
// If the element is not found the index is set to 0
func FindElem[T comparable](arr []T, elem T) (int, bool) {
	for i, val := range arr {
		if val == elem {
			return i, true
		}
	}
	return 0, false
}

// FindMatchRef will return the reference of the first element from 'arr' where 'match' is true
// it also returns a boolean that indicate wether the element is found or not
// if the element is not found, the pointer is nil
func FindMatchRef[T any](arr []T, match func(T) bool) (*T, bool) {
	for i, val := range arr {
		if match(val) {
			return &arr[i], true
		}
	}
	return nil, false
}

// FindMatch do the same thing than FindMatchRef except it return the index instead of a reference
// it return 0 if the element is not found
func FindMatch[T any](arr []T, match func(T) bool) (int, bool) {
	for i, val := range arr {
		if match(val) {
			return i, true
		}
	}
	return 0, false
}

// ReduceAs reduce the array 'arr' with the function reduce with possible different result type
func ReduceAs[T1, T2 any](arr []T1, reduce func(reduced T2, elem T1) T2) T2 {
	var result T2
	for _, elem := range arr {
		result = reduce(result, elem)
	}
	return result
}

// Reduce like ReduceAs except the result type is the same as the array type.
func Reduce[T any](arr []T, reduce func(reduced, elem T) T) T {
	return ReduceAs(arr, reduce)
}

// All will check if all element of 'arr' respect the condition 'check'
func All[T any](arr []T, check func(elem T) bool) bool {
	for _, value := range arr {
		if !check(value) {
			return false
		}
	}
	return true
}

// Compare two array with operator == on their element
// array of different size are not equal
func Compare[T comparable](arr1, arr2 []T) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for idx, elem := range arr1 {
		if elem != arr2[idx] {
			return false
		}
	}
	return true
}

// Compare two array with cmp function on their element
// array of different size are not equal
func CompareWith[T any](arr1, arr2 []T, cmp func(elem1, elem2 T) bool) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for idx, elem1 := range arr1 {
		if !cmp(elem1, arr2[idx]) {
			return false
		}
	}
	return true
}
