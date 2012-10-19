package builtin

import "fmt"

// Convert a builtin type to its autmented type.
func Augment(i interface{}) interface{} {
	switch i.(type) {
	case int:
		return Int(i.(int))
	case int8:
		return Int8(i.(int8))
	case int16:
		return Int16(i.(int16))
	case int32:
		return Int32(i.(int32))
	case int64:
		return Int64(i.(int64))
	case uint:
		return Uint(i.(uint))
	case uint8:
		return Uint8(i.(uint8))
	case uint16:
		return Uint16(i.(uint16))
	case uint32:
		return Uint32(i.(uint32))
	case uint64:
		return Uint64(i.(uint64))
	case float32:
		return Float32(i.(float32))
	case float64:
		return Float64(i.(float64))
	case string:
		return String(i.(string))
	}
	panic (fmt.Sprintf("%T is not a builtin", i))
}
