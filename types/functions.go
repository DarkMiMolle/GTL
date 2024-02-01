package types

import "reflect"

func Is[T any](val any) bool {
	_, ok := val.(T)
	return ok
}

func NewVal[T any](val T) *T {
	return &val
}

func Typeof[T any]() reflect.Type {
	var t T
	return reflect.TypeOf(t)
}
