package types

type IntegerTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type FloatTypes interface {
	~float32 | ~float64
}

type ComplexTypes interface {
	~complex64 | ~complex128
}

type NumberTypes interface {
	IntegerTypes | FloatTypes | ComplexTypes
}

type TextTypes interface {
	~string | ~rune
}

type BuiltinTypes interface {
	IntegerTypes | TextTypes | bool
}
