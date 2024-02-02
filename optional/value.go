package optional

import (
	"fmt"
	"reflect"
)

// Value[T] is an optional value of type T, meaning it can represent the lack of value.
type Value[T any] struct {
	value *T
}

// Value force to get the value of the optional value. If it hadn't a value it will panic
func (opt Value[T]) Value() T {
	return *opt.value
}

// ValueOr allow to get the containing value or a default value if the optional.Value didn't have any value
func (opt Value[T]) ValueOr(val T) T {
	if opt.value == nil {
		return val
	}
	return *opt.value
}

// HasValue return true if the optional.Value has a value
func (opt Value[T]) HasValue() bool {
	return opt.value != nil
}

// Set allows to set a value to the optional.Value
func (opt *Value[T]) Set(val T) {
	opt.value = &val
}

// SetNil allows to set "nil" to the optional.Value, meaning it remove the value from it
func (opt *Value[T]) SetNil() {
	opt.value = nil
}

// LookupValue return the containing value value or the a zero value if the optional.Value doesn't have a value, and a boolean to true if it contains a value
func (opt Value[T]) LookupValue() (T, bool) {
	var value T
	if opt.value == nil {
		return value, false
	}
	return *opt.value, true
}

// Eq is the "==" operator for optional.Value. It works on the sub type only.
func (opt Value[T]) Eq(val T) bool {
	if optVal := reflect.ValueOf(opt.ValueOr(val)); optVal.Comparable() {
		return opt.HasValue() && optVal.Equal(reflect.ValueOf(val)) // if opt hasValue then opt.ValueOf(val) will give the value of opt
	}
	return opt.HasValue() && reflect.DeepEqual(opt.Value(), val)
}

// EqQpt is the "==" operator for optional.Value between optional.Values
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

// String is the string representation of optional.Value
func (opt Value[T]) String() string {
	if opt.value == nil {
		return fmt.Sprintf("%v(null)", reflect.TypeOf(opt.value).Elem())
	}
	return fmt.Sprintf("%v", *opt.value)
}

// Some is a function to make an optional.Value with a value
func Some[T any](val T) Value[T] {
	return Value[T]{&val}
}

// Nil is a function to make an optional.Value without a value; it is equivalent to optional.Value's zero value
func Nil[T any]() Value[T] {
	return Missing[T]()
}

// Missing is equivalent to Nil but renamed for some usecase.
func Missing[T any]() Value[T] {
	return Value[T]{}
}

// Copy allows to make a deep copy of optional.Value. optional.Values share references by default.
func Copy[T any](opt Value[T]) Value[T] {
	if opt.value == nil {
		return opt
	}
	cpy := *opt.value
	return Value[T]{
		value: &cpy,
	}
}
