package array

func MapRef[T1, T2 any](arr []T1, action func(elem *T1) T2) []T2 {
	newArr := make([]T2, len(arr))
	for idx := range arr {
		newArr[idx] = action(&arr[idx])
	}
	return newArr
}

func Map[T1, T2 any](arr []T1, action func(elem T1) T2) []T2 {
	newArr := make([]T2, len(arr))
	for idx, value := range arr {
		newArr[idx] = action(value)
	}
	return newArr
}

func Reversed[T any](arr []T) []T {
	rev := make([]T, len(arr))
	for i := 0; i < len(arr)/2; i++ {
		rev[i] = arr[len(arr)-1-i]
		rev[len(arr)-1-i] = arr[i]
	}
	return rev
}

type Break bool

func ForEachReverse[T any](arr []T, action func(elem T) (breaking Break)) {
	for idx := range arr {
		if action(arr[len(arr)-1-idx]) {
			return
		}
	}
}

func Filter[T any](arr []T, filter func(elem T) bool) []T {
	filtered := make([]T, 0, len(arr))
	for _, val := range arr {
		if filter(val) {
			filtered = append(filtered, val)
		}
	}
	return filtered
}
