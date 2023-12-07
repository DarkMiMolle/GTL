package ints

import (
	"fmt"
	"reflect"
)

type BuiltinInts interface {
	int8 | int16 | int32 | int64 | int
}
type BuiltinUInts interface {
	uint8 | uint16 | uint32 | uint64 | uint
}

type BuilintDecimal interface {
	BuiltinInts | BuiltinUInts
}

type Hexadecimal []byte
type Hex = Hexadecimal

type filter byte

func (f filter) next() filter {
	if f == first {
		return second
	}
	return first
}
func (f filter) v() byte {
	return byte(f)
}

const (
	first  filter = 0b0000_1111
	second filter = 0b1111_0000
)

func HexValues[T BuilintDecimal](val T) Hexadecimal {
	hex := Hex{0}
	fill := first
	for val != 0 {
		current := &hex[0]
		elem := byte(val % 0x10)
		if fill == first {
			*current = elem
		} else if fill == second {
			elem = elem << 4
			*current = *current | elem
			hex = append(Hex{0}, hex...)
		}
		fill = fill.next()
		val /= 0x10
	}
	return hex
}
func (h Hexadecimal) String() string {
	hexStr := "0x"
	for _, hex := range h {
		hexStr += fmt.Sprintf("%x", hex)
	}
	return hexStr
}
func (h Hexadecimal) BitsSize() uintptr {
	return uintptr(len(h)) * 8
}
func genericDecimalCast[T BuilintDecimal](h Hex) T {
	var ret T
	h.checkSize(ret)
	for _, hex := range h {
		ret = ret << 8
		ret = ret | T(hex)
	}
	return ret
}
func (h Hexadecimal) checkSize(val any) {
	intSize := int(reflect.TypeOf(val).Size() * 8)
	if len(h)*8 > intSize {
		panic(fmt.Sprintf("%v can't be encoded with only %v bits. It needs at least %v bits", h, intSize, h.BitsSize()))
	}
}
func (h Hexadecimal) Int() int {
	return genericDecimalCast[int](h)
}
func (h Hexadecimal) Uint() uint {
	return genericDecimalCast[uint](h)
}
