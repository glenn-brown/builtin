// Package builtin supplies types with methods covering Go's builtin operators.
//
package builtin

type Int int

func (a Int) Eq(b interface{}) bool        { return a == b.(Int) }
func (a Int) NotEq(b interface{}) bool     { return a != b.(Int) }
func (a Int) Less(b interface{}) bool      { return a < b.(Int) }
func (a Int) LessEq(b interface{}) bool    { return a <= b.(Int) }
func (a Int) Greater(b interface{}) bool   { return a > b.(Int) }
func (a Int) GreaterEq(b interface{}) bool { return a >= b.(Int) }
func (a Int) Compare(b interface{}) int {
	c := b.(Int)
	if a <= c {
		if a < c {
			return -1
		}
		return 0
	}
	return 1
}
func (a Int) Max(b interface{}) interface{} {
	c := b.(Int)
	if c < a {
		c = a
	}
	return c
}
func (a Int) Min(b interface{}) interface{} {
	c := b.(Int)
	if c > a {
		c = a
	}
	return c
}
func (a Int) Add(b interface{}) interface{}        { return a + b.(Int) }
func (a Int) Sub(b interface{}) interface{}        { return a - b.(Int) }
func (a Int) Mul(b interface{}) interface{}        { return a * b.(Int) }
func (a Int) Div(b interface{}) interface{}        { return a / b.(Int) }
func (a Int) And(b interface{}) interface{}        { return a & b.(Int) }
func (a Int) Or(b interface{}) interface{}         { return a | b.(Int) }
func (a Int) Xor(b interface{}) interface{}        { return a ^ b.(Int) }
func (a Int) Not() interface{}                     { return ^a }
func (a Int) AndNot(b interface{}) interface{}     { return a &^ b.(Int) }
func (a Int) ShiftLeft(b interface{}) interface{}  { return a << b.(uint) }
func (a Int) ShiftRight(b interface{}) interface{} { return a >> b.(uint) }

type Int8 int8

func (a Int8) Eq(b interface{}) bool        { return a == b.(Int8) }
func (a Int8) NotEq(b interface{}) bool     { return a != b.(Int8) }
func (a Int8) Less(b interface{}) bool      { return a < b.(Int8) }
func (a Int8) LessEq(b interface{}) bool    { return a <= b.(Int8) }
func (a Int8) Greater(b interface{}) bool   { return a > b.(Int8) }
func (a Int8) GreaterEq(b interface{}) bool { return a >= b.(Int8) }
func (a Int8) Compare(b interface{}) int {
	c := b.(Int8)
	if a <= c {
		if a < c {
			return -1
		}
		return 0
	}
	return 1
}
func (a Int8) Max(b interface{}) interface{} {
	c := b.(Int8)
	if c < a {
		c = a
	}
	return c
}
func (a Int8) Min(b interface{}) interface{} {
	c := b.(Int8)
	if c > a {
		c = a
	}
	return c
}
func (a Int8) Add(b interface{}) interface{}        { return a + b.(Int8) }
func (a Int8) Sub(b interface{}) interface{}        { return a - b.(Int8) }
func (a Int8) Mul(b interface{}) interface{}        { return a * b.(Int8) }
func (a Int8) Div(b interface{}) interface{}        { return a / b.(Int8) }
func (a Int8) And(b interface{}) interface{}        { return a & b.(Int8) }
func (a Int8) Or(b interface{}) interface{}         { return a | b.(Int8) }
func (a Int8) Xor(b interface{}) interface{}        { return a ^ b.(Int8) }
func (a Int8) Not() interface{}                     { return ^a }
func (a Int8) AndNot(b interface{}) interface{}     { return a &^ b.(Int8) }
func (a Int8) ShiftLeft(b interface{}) interface{}  { return a << b.(uint) }
func (a Int8) ShiftRight(b interface{}) interface{} { return a >> b.(uint) }

type Int16 int16

func (a Int16) Eq(b interface{}) bool        { return a == b.(Int16) }
func (a Int16) NotEq(b interface{}) bool     { return a != b.(Int16) }
func (a Int16) Less(b interface{}) bool      { return a < b.(Int16) }
func (a Int16) LessEq(b interface{}) bool    { return a <= b.(Int16) }
func (a Int16) Greater(b interface{}) bool   { return a > b.(Int16) }
func (a Int16) GreaterEq(b interface{}) bool { return a >= b.(Int16) }
func (a Int16) Compare(b interface{}) int {
	c := b.(Int16)
	if a <= c {
		if a < c {
			return -1
		}
		return 0
	}
	return 1
}
func (a Int16) Max(b interface{}) interface{} {
	c := b.(Int16)
	if c < a {
		c = a
	}
	return c
}
func (a Int16) Min(b interface{}) interface{} {
	c := b.(Int16)
	if c > a {
		c = a
	}
	return c
}
func (a Int16) Add(b interface{}) interface{}        { return a + b.(Int16) }
func (a Int16) Sub(b interface{}) interface{}        { return a - b.(Int16) }
func (a Int16) Mul(b interface{}) interface{}        { return a * b.(Int16) }
func (a Int16) Div(b interface{}) interface{}        { return a / b.(Int16) }
func (a Int16) And(b interface{}) interface{}        { return a & b.(Int16) }
func (a Int16) Or(b interface{}) interface{}         { return a | b.(Int16) }
func (a Int16) Xor(b interface{}) interface{}        { return a ^ b.(Int16) }
func (a Int16) Not() interface{}                     { return ^a }
func (a Int16) AndNot(b interface{}) interface{}     { return a &^ b.(Int16) }
func (a Int16) ShiftLeft(b interface{}) interface{}  { return a << b.(uint) }
func (a Int16) ShiftRight(b interface{}) interface{} { return a >> b.(uint) }

type Int32 int32

func (a Int32) Eq(b interface{}) bool        { return a == b.(Int32) }
func (a Int32) NotEq(b interface{}) bool     { return a != b.(Int32) }
func (a Int32) Less(b interface{}) bool      { return a < b.(Int32) }
func (a Int32) LessEq(b interface{}) bool    { return a <= b.(Int32) }
func (a Int32) Greater(b interface{}) bool   { return a > b.(Int32) }
func (a Int32) GreaterEq(b interface{}) bool { return a >= b.(Int32) }
func (a Int32) Compare(b interface{}) int {
	c := b.(Int32)
	if a <= c {
		if a < c {
			return -1
		}
		return 0
	}
	return 1
}
func (a Int32) Max(b interface{}) interface{} {
	c := b.(Int32)
	if c < a {
		c = a
	}
	return c
}
func (a Int32) Min(b interface{}) interface{} {
	c := b.(Int32)
	if c > a {
		c = a
	}
	return c
}
func (a Int32) Add(b interface{}) interface{}        { return a + b.(Int32) }
func (a Int32) Sub(b interface{}) interface{}        { return a - b.(Int32) }
func (a Int32) Mul(b interface{}) interface{}        { return a * b.(Int32) }
func (a Int32) Div(b interface{}) interface{}        { return a / b.(Int32) }
func (a Int32) And(b interface{}) interface{}        { return a & b.(Int32) }
func (a Int32) Or(b interface{}) interface{}         { return a | b.(Int32) }
func (a Int32) Xor(b interface{}) interface{}        { return a ^ b.(Int32) }
func (a Int32) Not() interface{}                     { return ^a }
func (a Int32) AndNot(b interface{}) interface{}     { return a &^ b.(Int32) }
func (a Int32) ShiftLeft(b interface{}) interface{}  { return a << b.(uint) }
func (a Int32) ShiftRight(b interface{}) interface{} { return a >> b.(uint) }

type Int64 int64

func (a Int64) Eq(b interface{}) bool        { return a == b.(Int64) }
func (a Int64) NotEq(b interface{}) bool     { return a != b.(Int64) }
func (a Int64) Less(b interface{}) bool      { return a < b.(Int64) }
func (a Int64) LessEq(b interface{}) bool    { return a <= b.(Int64) }
func (a Int64) Greater(b interface{}) bool   { return a > b.(Int64) }
func (a Int64) GreaterEq(b interface{}) bool { return a >= b.(Int64) }
func (a Int64) Compare(b interface{}) int {
	c := b.(Int64)
	if a <= c {
		if a < c {
			return -1
		}
		return 0
	}
	return 1
}
func (a Int64) Max(b interface{}) interface{} {
	c := b.(Int64)
	if c < a {
		c = a
	}
	return c
}
func (a Int64) Min(b interface{}) interface{} {
	c := b.(Int64)
	if c > a {
		c = a
	}
	return c
}
func (a Int64) Add(b interface{}) interface{}        { return a + b.(Int64) }
func (a Int64) Sub(b interface{}) interface{}        { return a - b.(Int64) }
func (a Int64) Mul(b interface{}) interface{}        { return a * b.(Int64) }
func (a Int64) Div(b interface{}) interface{}        { return a / b.(Int64) }
func (a Int64) And(b interface{}) interface{}        { return a & b.(Int64) }
func (a Int64) Or(b interface{}) interface{}         { return a | b.(Int64) }
func (a Int64) Xor(b interface{}) interface{}        { return a ^ b.(Int64) }
func (a Int64) Not() interface{}                     { return ^a }
func (a Int64) AndNot(b interface{}) interface{}     { return a &^ b.(Int64) }
func (a Int64) ShiftLeft(b interface{}) interface{}  { return a << b.(uint) }
func (a Int64) ShiftRight(b interface{}) interface{} { return a >> b.(uint) }

type Uint uint

func (a Uint) Eq(b interface{}) bool        { return a == b.(Uint) }
func (a Uint) NotEq(b interface{}) bool     { return a != b.(Uint) }
func (a Uint) Less(b interface{}) bool      { return a < b.(Uint) }
func (a Uint) LessEq(b interface{}) bool    { return a <= b.(Uint) }
func (a Uint) Greater(b interface{}) bool   { return a > b.(Uint) }
func (a Uint) GreaterEq(b interface{}) bool { return a >= b.(Uint) }
func (a Uint) Compare(b interface{}) int {
	c := b.(Uint)
	if a <= c {
		if a < c {
			return -1
		}
		return 0
	}
	return 1
}
func (a Uint) Max(b interface{}) interface{} {
	c := b.(Uint)
	if c < a {
		c = a
	}
	return c
}
func (a Uint) Min(b interface{}) interface{} {
	c := b.(Uint)
	if c > a {
		c = a
	}
	return c
}
func (a Uint) Add(b interface{}) interface{}        { return a + b.(Uint) }
func (a Uint) Sub(b interface{}) interface{}        { return a - b.(Uint) }
func (a Uint) Mul(b interface{}) interface{}        { return a * b.(Uint) }
func (a Uint) Div(b interface{}) interface{}        { return a / b.(Uint) }
func (a Uint) And(b interface{}) interface{}        { return a & b.(Uint) }
func (a Uint) Or(b interface{}) interface{}         { return a | b.(Uint) }
func (a Uint) Xor(b interface{}) interface{}        { return a ^ b.(Uint) }
func (a Uint) Not() interface{}                     { return ^a }
func (a Uint) AndNot(b interface{}) interface{}     { return a &^ b.(Uint) }
func (a Uint) ShiftLeft(b interface{}) interface{}  { return a << b.(uint) }
func (a Uint) ShiftRight(b interface{}) interface{} { return a >> b.(uint) }

type Uint8 uint8

func (a Uint8) Eq(b interface{}) bool        { return a == b.(Uint8) }
func (a Uint8) NotEq(b interface{}) bool     { return a != b.(Uint8) }
func (a Uint8) Less(b interface{}) bool      { return a < b.(Uint8) }
func (a Uint8) LessEq(b interface{}) bool    { return a <= b.(Uint8) }
func (a Uint8) Greater(b interface{}) bool   { return a > b.(Uint8) }
func (a Uint8) GreaterEq(b interface{}) bool { return a >= b.(Uint8) }
func (a Uint8) Compare(b interface{}) int {
	c := b.(Uint8)
	if a <= c {
		if a < c {
			return -1
		}
		return 0
	}
	return 1
}
func (a Uint8) Max(b interface{}) interface{} {
	c := b.(Uint8)
	if c < a {
		c = a
	}
	return c
}
func (a Uint8) Min(b interface{}) interface{} {
	c := b.(Uint8)
	if c > a {
		c = a
	}
	return c
}
func (a Uint8) Add(b interface{}) interface{}        { return a + b.(Uint8) }
func (a Uint8) Sub(b interface{}) interface{}        { return a - b.(Uint8) }
func (a Uint8) Mul(b interface{}) interface{}        { return a * b.(Uint8) }
func (a Uint8) Div(b interface{}) interface{}        { return a / b.(Uint8) }
func (a Uint8) And(b interface{}) interface{}        { return a & b.(Uint8) }
func (a Uint8) Or(b interface{}) interface{}         { return a | b.(Uint8) }
func (a Uint8) Xor(b interface{}) interface{}        { return a ^ b.(Uint8) }
func (a Uint8) Not() interface{}                     { return ^a }
func (a Uint8) AndNot(b interface{}) interface{}     { return a &^ b.(Uint8) }
func (a Uint8) ShiftLeft(b interface{}) interface{}  { return a << b.(uint) }
func (a Uint8) ShiftRight(b interface{}) interface{} { return a >> b.(uint) }

type Uint16 uint16

func (a Uint16) Eq(b interface{}) bool        { return a == b.(Uint16) }
func (a Uint16) NotEq(b interface{}) bool     { return a != b.(Uint16) }
func (a Uint16) Less(b interface{}) bool      { return a < b.(Uint16) }
func (a Uint16) LessEq(b interface{}) bool    { return a <= b.(Uint16) }
func (a Uint16) Greater(b interface{}) bool   { return a > b.(Uint16) }
func (a Uint16) GreaterEq(b interface{}) bool { return a >= b.(Uint16) }
func (a Uint16) Compare(b interface{}) int {
	c := b.(Uint16)
	if a <= c {
		if a < c {
			return -1
		}
		return 0
	}
	return 1
}
func (a Uint16) Max(b interface{}) interface{} {
	c := b.(Uint16)
	if c < a {
		c = a
	}
	return c
}
func (a Uint16) Min(b interface{}) interface{} {
	c := b.(Uint16)
	if c > a {
		c = a
	}
	return c
}
func (a Uint16) Add(b interface{}) interface{}        { return a + b.(Uint16) }
func (a Uint16) Sub(b interface{}) interface{}        { return a - b.(Uint16) }
func (a Uint16) Mul(b interface{}) interface{}        { return a * b.(Uint16) }
func (a Uint16) Div(b interface{}) interface{}        { return a / b.(Uint16) }
func (a Uint16) And(b interface{}) interface{}        { return a & b.(Uint16) }
func (a Uint16) Or(b interface{}) interface{}         { return a | b.(Uint16) }
func (a Uint16) Xor(b interface{}) interface{}        { return a ^ b.(Uint16) }
func (a Uint16) Not() interface{}                     { return ^a }
func (a Uint16) AndNot(b interface{}) interface{}     { return a &^ b.(Uint16) }
func (a Uint16) ShiftLeft(b interface{}) interface{}  { return a << b.(uint) }
func (a Uint16) ShiftRight(b interface{}) interface{} { return a >> b.(uint) }

type Uint32 uint32

func (a Uint32) Eq(b interface{}) bool        { return a == b.(Uint32) }
func (a Uint32) NotEq(b interface{}) bool     { return a != b.(Uint32) }
func (a Uint32) Less(b interface{}) bool      { return a < b.(Uint32) }
func (a Uint32) LessEq(b interface{}) bool    { return a <= b.(Uint32) }
func (a Uint32) Greater(b interface{}) bool   { return a > b.(Uint32) }
func (a Uint32) GreaterEq(b interface{}) bool { return a >= b.(Uint32) }
func (a Uint32) Compare(b interface{}) int {
	c := b.(Uint32)
	if a <= c {
		if a < c {
			return -1
		}
		return 0
	}
	return 1
}
func (a Uint32) Max(b interface{}) interface{} {
	c := b.(Uint32)
	if c < a {
		c = a
	}
	return c
}
func (a Uint32) Min(b interface{}) interface{} {
	c := b.(Uint32)
	if c > a {
		c = a
	}
	return c
}
func (a Uint32) Add(b interface{}) interface{}        { return a + b.(Uint32) }
func (a Uint32) Sub(b interface{}) interface{}        { return a - b.(Uint32) }
func (a Uint32) Mul(b interface{}) interface{}        { return a * b.(Uint32) }
func (a Uint32) Div(b interface{}) interface{}        { return a / b.(Uint32) }
func (a Uint32) And(b interface{}) interface{}        { return a & b.(Uint32) }
func (a Uint32) Or(b interface{}) interface{}         { return a | b.(Uint32) }
func (a Uint32) Xor(b interface{}) interface{}        { return a ^ b.(Uint32) }
func (a Uint32) Not() interface{}                     { return ^a }
func (a Uint32) AndNot(b interface{}) interface{}     { return a &^ b.(Uint32) }
func (a Uint32) ShiftLeft(b interface{}) interface{}  { return a << b.(uint) }
func (a Uint32) ShiftRight(b interface{}) interface{} { return a >> b.(uint) }

type Uint64 uint64

func (a Uint64) Eq(b interface{}) bool        { return a == b.(Uint64) }
func (a Uint64) NotEq(b interface{}) bool     { return a != b.(Uint64) }
func (a Uint64) Less(b interface{}) bool      { return a < b.(Uint64) }
func (a Uint64) LessEq(b interface{}) bool    { return a <= b.(Uint64) }
func (a Uint64) Greater(b interface{}) bool   { return a > b.(Uint64) }
func (a Uint64) GreaterEq(b interface{}) bool { return a >= b.(Uint64) }
func (a Uint64) Compare(b interface{}) int {
	c := b.(Uint64)
	if a <= c {
		if a < c {
			return -1
		}
		return 0
	}
	return 1
}
func (a Uint64) Max(b interface{}) interface{} {
	c := b.(Uint64)
	if c < a {
		c = a
	}
	return c
}
func (a Uint64) Min(b interface{}) interface{} {
	c := b.(Uint64)
	if c > a {
		c = a
	}
	return c
}
func (a Uint64) Add(b interface{}) interface{}        { return a + b.(Uint64) }
func (a Uint64) Sub(b interface{}) interface{}        { return a - b.(Uint64) }
func (a Uint64) Mul(b interface{}) interface{}        { return a * b.(Uint64) }
func (a Uint64) Div(b interface{}) interface{}        { return a / b.(Uint64) }
func (a Uint64) And(b interface{}) interface{}        { return a & b.(Uint64) }
func (a Uint64) Or(b interface{}) interface{}         { return a | b.(Uint64) }
func (a Uint64) Xor(b interface{}) interface{}        { return a ^ b.(Uint64) }
func (a Uint64) Not() interface{}                     { return ^a }
func (a Uint64) AndNot(b interface{}) interface{}     { return a &^ b.(Uint64) }
func (a Uint64) ShiftLeft(b interface{}) interface{}  { return a << b.(uint) }
func (a Uint64) ShiftRight(b interface{}) interface{} { return a >> b.(uint) }

type Uintptr uintptr

func (a Uintptr) Eq(b interface{}) bool        { return a == b.(Uintptr) }
func (a Uintptr) NotEq(b interface{}) bool     { return a != b.(Uintptr) }
func (a Uintptr) Less(b interface{}) bool      { return a < b.(Uintptr) }
func (a Uintptr) LessEq(b interface{}) bool    { return a <= b.(Uintptr) }
func (a Uintptr) Greater(b interface{}) bool   { return a > b.(Uintptr) }
func (a Uintptr) GreaterEq(b interface{}) bool { return a >= b.(Uintptr) }
func (a Uintptr) Compare(b interface{}) int {
	c := b.(Uintptr)
	if a <= c {
		if a < c {
			return -1
		}
		return 0
	}
	return 1
}
func (a Uintptr) Max(b interface{}) interface{} {
	c := b.(Uintptr)
	if c < a {
		c = a
	}
	return c
}
func (a Uintptr) Min(b interface{}) interface{} {
	c := b.(Uintptr)
	if c > a {
		c = a
	}
	return c
}
func (a Uintptr) Add(b interface{}) interface{}        { return a + b.(Uintptr) }
func (a Uintptr) Sub(b interface{}) interface{}        { return a - b.(Uintptr) }
func (a Uintptr) Mul(b interface{}) interface{}        { return a * b.(Uintptr) }
func (a Uintptr) Div(b interface{}) interface{}        { return a / b.(Uintptr) }
func (a Uintptr) And(b interface{}) interface{}        { return a & b.(Uintptr) }
func (a Uintptr) Or(b interface{}) interface{}         { return a | b.(Uintptr) }
func (a Uintptr) Xor(b interface{}) interface{}        { return a ^ b.(Uintptr) }
func (a Uintptr) Not() interface{}                     { return ^a }
func (a Uintptr) AndNot(b interface{}) interface{}     { return a &^ b.(Uintptr) }
func (a Uintptr) ShiftLeft(b interface{}) interface{}  { return a << b.(uint) }
func (a Uintptr) ShiftRight(b interface{}) interface{} { return a >> b.(uint) }

type Float64 float64

func (a Float64) Eq(b interface{}) bool        { return a == b.(Float64) }
func (a Float64) NotEq(b interface{}) bool     { return a != b.(Float64) }
func (a Float64) Less(b interface{}) bool      { return a < b.(Float64) }
func (a Float64) LessEq(b interface{}) bool    { return a <= b.(Float64) }
func (a Float64) Greater(b interface{}) bool   { return a > b.(Float64) }
func (a Float64) GreaterEq(b interface{}) bool { return a >= b.(Float64) }
func (a Float64) Compare(b interface{}) int {
	c := b.(Float64)
	if a <= c {
		if a < c {
			return -1
		}
		return 0
	}
	return 1
}
func (a Float64) Max(b interface{}) interface{} {
	c := b.(Float64)
	if c < a {
		c = a
	}
	return c
}
func (a Float64) Min(b interface{}) interface{} {
	c := b.(Float64)
	if c > a {
		c = a
	}
	return c
}
func (a Float64) Add(b interface{}) interface{} { return a + b.(Float64) }
func (a Float64) Sub(b interface{}) interface{} { return a - b.(Float64) }
func (a Float64) Mul(b interface{}) interface{} { return a * b.(Float64) }
func (a Float64) Div(b interface{}) interface{} { return a / b.(Float64) }

type Float32 float32

func (a Float32) Eq(b interface{}) bool        { return a == b.(Float32) }
func (a Float32) NotEq(b interface{}) bool     { return a != b.(Float32) }
func (a Float32) Less(b interface{}) bool      { return a < b.(Float32) }
func (a Float32) LessEq(b interface{}) bool    { return a <= b.(Float32) }
func (a Float32) Greater(b interface{}) bool   { return a > b.(Float32) }
func (a Float32) GreaterEq(b interface{}) bool { return a >= b.(Float32) }
func (a Float32) Compare(b interface{}) int {
	c := b.(Float32)
	if a <= c {
		if a < c {
			return -1
		}
		return 0
	}
	return 1
}
func (a Float32) Max(b interface{}) interface{} {
	c := b.(Float32)
	if c < a {
		c = a
	}
	return c
}
func (a Float32) Min(b interface{}) interface{} {
	c := b.(Float32)
	if c > a {
		c = a
	}
	return c
}
func (a Float32) Add(b interface{}) interface{} { return a + b.(Float32) }
func (a Float32) Sub(b interface{}) interface{} { return a - b.(Float32) }
func (a Float32) Mul(b interface{}) interface{} { return a * b.(Float32) }
func (a Float32) Div(b interface{}) interface{} { return a / b.(Float32) }

type String string

func (a String) Eq(b interface{}) bool        { return a == b.(String) }
func (a String) NotEq(b interface{}) bool     { return a != b.(String) }
func (a String) Less(b interface{}) bool      { return a < b.(String) }
func (a String) LessEq(b interface{}) bool    { return a <= b.(String) }
func (a String) Greater(b interface{}) bool   { return a > b.(String) }
func (a String) GreaterEq(b interface{}) bool { return a >= b.(String) }
func (a String) Compare(b interface{}) int {
	c := b.(String)
	if a <= c {
		if a < c {
			return -1
		}
		return 0
	}
	return 1
}
func (a String) Max(b interface{}) interface{} {
	c := b.(String)
	if c < a {
		c = a
	}
	return c
}
func (a String) Min(b interface{}) interface{} {
	c := b.(String)
	if c > a {
		c = a
	}
	return c
}
