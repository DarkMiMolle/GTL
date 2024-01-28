package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapRef(t *testing.T) {
	t.Run("to string", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := MapRef(arr, func(i *int) string { return fmt.Sprint(*i) })
		expect := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("MapRef() = %v, want %v", got, expect)
		}
	})

	t.Run("get address", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := MapRef(arr, func(i *int) *int { return i })
		expect := []*int{&arr[0], &arr[1], &arr[2], &arr[3], &arr[4], &arr[5], &arr[6], &arr[7], &arr[8], &arr[9]}
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("MapRef() = %v, want %v", got, expect)
		}
	})
}

func TestMap(t *testing.T) {
	t.Run("to string", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := Map(arr, func(i int) string { return fmt.Sprint(i) })
		expect := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("Map() = %v, want %v", got, expect)
		}
	})
}

func TestReversed(t *testing.T) {
	t.Run("reverse odd length array", func(t *testing.T) {
		arr := []int{1, 2, 3, 4}
		got := Reversed(arr)
		expect := []int{4, 3, 2, 1}

		if !reflect.DeepEqual(got, expect) {
			t.Errorf("Reversed(%v) = %v, want %v", arr, got, expect)
		}
	})
}
