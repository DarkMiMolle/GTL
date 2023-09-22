package optional

import (
	"reflect"
)

type Value[T any] struct {
	value *T
}

func (opt Value[T]) Value() T {
	return *opt.value
}
func (opt Value[T]) ValueOr(val T) T {
	if opt.value == nil {
		return val
	}
	return *opt.value
}
func (opt Value[T]) HasValue() bool {
	return opt.value != nil
}
func (opt *Value[T]) Set(val T) {
	opt.value = &val
}
func (opt *Value[T]) SetNil() {
	opt.value = nil
}
func (opt Value[T]) LookupValue() (T, bool) {
	var value T
	if opt.value == nil {
		return value, false
	}
	return *opt.value, true
}
func (opt Value[T]) Eq(val T) bool {
	if optVal := reflect.ValueOf(opt.ValueOr(val)); optVal.Comparable() {
		return opt.HasValue() && optVal.Equal(reflect.ValueOf(val)) // if opt hasValue then opt.ValueOf(val) will give the value of opt
	}
	return opt.HasValue() && reflect.DeepEqual(opt.Value(), val)
}
func (opt Value[T]) EqOpt(val Value[T]) bool {
	if opt.HasValue() != val.HasValue() {
		return false
	}
	// from here we know that if one has a value, both have
	// and if one has nil, both have
	var zero T
	/*
		if they have nil:
			opt.Eq(val.ValueOr(zero)) => false
			val.Eq(opt.ValueOr(zero)) => false
			false == false --> they are equal
		else:
			opt.Eq(val.ValueOr(zero)) <=> opt.Value() == val.Value()
			val.Eq(opt.ValueOr(zero)) <=> val.Value() == opt.Value()
		 	opt.Value() == val.Value() <=> val.Value() == opt.Value()
			--> if they are the same value it will be true, and false otherwise
	*/
	return opt.Eq(val.ValueOr(zero)) == val.Eq(opt.ValueOr(zero))
}

func Some[T any](val T) Value[T] {
	return Value[T]{&val}
}
func Missing[T any]() Value[T] {
	return Value[T]{}
}
func Nil[T any]() Value[T] {
	return Missing[T]()
}
