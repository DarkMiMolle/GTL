package optional

import (
	"fmt"
	"reflect"
)

func Access[OBJ, F any](opt Value[OBJ], field string) Value[F] {
	if !opt.HasValue() {
		return Nil[F]()
	}
	var f F
	expectedFieldType := reflect.TypeOf(f)

	value := reflect.ValueOf(opt.value).Elem()
	if field, fieldExists := value.Type().FieldByName(field); fieldExists {
		if field.Type != expectedFieldType {
			return Nil[F]()
		}
		field := value.FieldByName(field.Name)
		if field.CanAddr() {
			return Value[F]{value: field.Addr().Interface().(*F)}
		}
		return Some[F](field.Interface().(F))
	}
	return Nil[F]()
}

func TryExpr[T any](action func() T) (ret Value[T]) {
	defer func() {
		if recover() != nil {
			ret = Nil[T]()
		}
	}()
	return Some(action())
}

func TryErr[T any](action func() T) (ret Value[T], err error) {
	defer func() {
		if rec := recover(); rec != nil {
			switch rec := rec.(type) {
			case error:
				err = rec
			default:
				err = fmt.Errorf("%v", rec)
			}
			ret = Nil[T]()
		}
	}()
	return Some(action()), nil
}

func Try(action func()) error {
	_, err := TryErr[struct{}](func() (r struct{}) { action(); return })
	return err
}
