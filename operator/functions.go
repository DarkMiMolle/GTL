package operator

import (
	"github.com/DarkMiMolle/GTL/types"
	"reflect"
)

func BuiltinPlus[T interface {
	types.NumberTypes | types.TextTypes
}](left, right T) T {
	return left + right
}

func BuiltinMinus[T interface{ types.NumberTypes }](left, right T) T {
	return left - right
}

func BuiltinMult[T interface{ types.NumberTypes }](left, right T) T {
	return left * right
}

func BuiltinDiv[T interface{ types.NumberTypes }](left, right T) T {
	return left / right
}

func BuiltinMod[T interface{ types.IntegerTypes }](left, right T) T {
	return left % right
}

func operatorWrapper(f any) func(any, any) any {
	return func(left, right any) any {
		return reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(left), reflect.ValueOf(right)})[0].Interface()
	}
}

var plusOperator map[reflect.Kind]func(any, any) any

func innerStructPlus(left, right any) any {
	leftValue := reflect.ValueOf(left)
	rightValue := reflect.ValueOf(right)
	returnValue := reflect.New(leftValue.Type())

	for i := 0; i < leftValue.NumField(); i++ {
		leftField := leftValue.Field(i)
		rightField := rightValue.Field(i)
		returnField := returnValue.Elem().Field(i).Addr()
		plus, exists := plusOperator[leftField.Kind()]
		if !exists {
			continue
		}
		returnField.Elem().Set(reflect.ValueOf(plus(leftField.Interface(), rightField.Interface())))
	}
	return returnValue.Elem().Interface()
}

func init() {
	plusOperator = map[reflect.Kind]func(any, any) any{
		reflect.Float32:    operatorWrapper(BuiltinPlus[float32]),
		reflect.Float64:    operatorWrapper(BuiltinPlus[float64]),
		reflect.Int:        operatorWrapper(BuiltinPlus[int]),
		reflect.Int16:      operatorWrapper(BuiltinPlus[int16]),
		reflect.Int32:      operatorWrapper(BuiltinPlus[int32]),
		reflect.Int64:      operatorWrapper(BuiltinPlus[int64]),
		reflect.Int8:       operatorWrapper(BuiltinPlus[int8]),
		reflect.Uint:       operatorWrapper(BuiltinPlus[uint]),
		reflect.Uint8:      operatorWrapper(BuiltinPlus[uint8]),
		reflect.Uint16:     operatorWrapper(BuiltinPlus[uint16]),
		reflect.Uint32:     operatorWrapper(BuiltinPlus[uint32]),
		reflect.Uint64:     operatorWrapper(BuiltinPlus[uint64]),
		reflect.String:     operatorWrapper(BuiltinPlus[string]),
		reflect.Complex64:  operatorWrapper(BuiltinPlus[complex64]),
		reflect.Complex128: operatorWrapper(BuiltinPlus[complex128]),
		reflect.Struct:     innerStructPlus,
	}
}

func StructPlus[T any](left, right T) T {
	if reflect.TypeOf(left).Kind() != reflect.Struct {
		panic("StructPlus must be called with structure type")
	}
	return innerStructPlus(left, right).(T)
}

// func StructMinus[T any](left, right T) T {

// }
// func StructMult[T any](left, right T) T {

// }
// func StructDiv[T any](left, right T) T {

// }
// func StructMod[T any](left, right T) T {

// }
