#!/bin/sh

# Replace "@" with T in input stream
at () {
    local T="$1"
    sed -e "s/@/$T/g"
}

# Type classes
#
# These functions emit per-type-class methods

Logical () {
    at "$1" <<EOF
func (a @) And(b interface{}) interface{} 	{ return a & b.(@) }
func (a @) Or(b interface{})  interface{}	{ return a | b.(@) }
func (a @) Xor(b interface{}) interface{}	{ return a ^ b.(@) }
func (a @) Not() interface{}			{ return ^a }
func (a @) AndNot(b interface{}) interface{}	{ return a &^ b.(@) }
EOF
}

Bits () {
    Logical $1
    at $1 <<EOF
func (a @) ShiftLeft(b interface{}) interface{}	{ return a << b.(uint) }
func (a @) ShiftRight(b interface{}) interface{} { return a >> b.(uint) }
EOF
}

Eq () {
    at "$1" <<EOF
func (a @) Eq(b interface{}) bool		{ return a == b.(@) }
func (a @) NotEq(b interface{}) bool		{ return a != b.(@) }
EOF
}

Ordinal () {
    Eq "$1"
    at "$1" <<EOF
func (a @) Less(b interface{}) bool		{ return a < b.(@) }
func (a @) LessEq(b interface{}) bool		{ return a <= b.(@) }
func (a @) Greater(b interface{}) bool		{ return a > b.(@) }
func (a @) GreaterEq(b interface{}) bool	{ return a >= b.(@) }
func (a @) Compare(b interface{}) int		{ c:=b.(@); if a <= c { if a < c { return -1 } ; return 0 } ; return 1 }
func (a @) Max(b interface{}) interface{}	{ c:=b.(@); if c < a { c = a }; return c }
func (a @) Min(b interface{}) interface{}	{ c:=b.(@); if c > a { c = a }; return c }
EOF
}
    
Numerical () {
    Ordinal "$1"
    at "$1" <<EOF
func (a @) Add(b interface{}) interface{}	{ return a + b.(@) }
func (a @) Sub(b interface{}) interface{}	{ return a - b.(@) }
func (a @) Mul(b interface{}) interface{}	{ return a * b.(@) }
func (a @) Div(b interface{}) interface{}	{ return a / b.(@) }
EOF
}

Integral () {
    at "$1" <<EOF
func (a @) Mod(b interface{}) interface{}	{ return a % b.(@) }
EOF
}

# Emit definitions for a type.
# 
type () {
    local T="$1" ; shift
    local t="$1" ; shift
    # Remaining arguments are nonredundant type classes.
    echo "type $T $t"
    for class in "$@" ; do
	$class $T
    done
}

builtin () {
    cat <<EOF
// Package builtin supplies types with methods covering Go's builtin operators.
//
package builtin
EOF
    type Int int		Numerical Bits
    type Int8 int8		Numerical Bits
    type Int16 int16		Numerical Bits
    type Int32 int32		Numerical Bits
    type Int64 int64		Numerical Bits
    type Uint uint		Numerical Bits
    type Uint8 uint8		Numerical Bits
    type Uint16 uint16		Numerical Bits
    type Uint32 uint32		Numerical Bits
    type Uint64 uint64		Numerical Bits
    type Float64 float64	Numerical
    type Float32 float32	Numerical
    type String string		Ordinal
}

builtin > builtin.go
gofmt -w builtin.go
