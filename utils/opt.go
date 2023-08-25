package utils

import "reflect"

// TODO: test

type Opt[T any] struct {
	value *T
}

func (opt *Opt[T]) Value() T {
	return *opt.value
}
func (opt *Opt[T]) ValueOr(val T) T {
	if opt.value == nil {
		return val
	}
	return *opt.value
}
func (opt *Opt[T]) HasValue() bool {
	return opt.value != nil
}
func (opt *Opt[T]) Set(val T) {
	opt.value = &val
}
func (opt *Opt[T]) SetNil() {
	opt.value = nil
}
func (opt *Opt[T]) LookupValue() (T, bool) {
	var value T
	if opt.value == nil {
		return value, false
	}
	return *opt.value, true
}
func (opt *Opt[T]) Eq(val T) bool {
	if optVal := reflect.ValueOf(opt.ValueOr(val)); optVal.Comparable() {
		return opt.HasValue() && optVal.Equal(reflect.ValueOf(val)) // if opt hasValue then opt.ValueOf(val) will give the value of opt
	}
	return opt.HasValue() && reflect.DeepEqual(opt.Value(), val)
}
func (opt *Opt[T]) EqOpt(val Opt[T]) bool {
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
