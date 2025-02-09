package delve

import "reflect"

type Numeric interface {
	float64 | float32 | int | int64 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8
}

func AnyToNumeric[T Numeric](num any) (val T, ok bool) {
	ok = true
	switch casted := num.(type) {
	case T:
		return casted, true
	case float64:
		val = T(casted)
	case float32:
		val = T(casted)
	case int:
		val = T(casted)
	case int64:
		val = T(casted)
	case int32:
		val = T(casted)
	case int16:
		val = T(casted)
	case int8:
		val = T(casted)
	case uint:
		val = T(casted)
	case uint64:
		val = T(casted)
	case uint32:
		val = T(casted)
	case uint16:
		val = T(casted)
	case uint8:
		val = T(casted)
	default:
		ok = false
	}
	return
}

func getTyped[T any](fm *Navigator, qual IQual, _default ...T) T {
	var defaultVal T
	if len(_default) > 0 {
		defaultVal = _default[0]
	}
	if val, ok := fm.QualGet(qual); ok {
		if casted, ok := val.(T); ok {
			return casted
		}
	}
	return defaultVal
}

// Len attempts to get the length of a value associated with a given qualifier.
//
// It retrieves the value associated with `qual`. If found, and the value is one of
// the following types: Chan, Map, Array, Slice, or String, the function returns
// the length of the value. Otherwise (the qualifier is not found, or the value
// has a type that doesn't have a defined length concept in Go), it returns -1.
func (fm *Navigator) Len(qual IQual) int {
	val, ok := fm.QualGet(qual)
	if !ok {
		return -1
	}
	refVal := reflect.ValueOf(val)
	switch refVal.Kind() {
	case reflect.Chan, reflect.Map, reflect.Array, reflect.Slice, reflect.String:
		return refVal.Len()
	}
	return -1
}

// SafeInterface retrieves a value associated with the given qualifier `qual`.
// It returns the value if found and if it's type-assignable to the type of `defaultVal`.
// If the qualifier is not found or the type is not assignable, it returns `defaultVal`.
func (fm *Navigator) SafeInterface(qual IQual, defaultVal any) any {
	val, ok := fm.QualGet(qual)
	if !ok {
		return defaultVal
	}
	if reflect.TypeOf(defaultVal).AssignableTo(reflect.TypeOf(val)) {
		return val
	}
	return defaultVal
}

// Get interface or default.
func (fm *Navigator) Interface(qual IQual, _default ...any) any {
	return getTyped(fm, qual, _default...)
}

// Get string or default
func (fm *Navigator) String(qual IQual, _default ...string) string {
	return getTyped(fm, qual, _default...)
}

// Get boolean or default
func (fm *Navigator) Bool(qual IQual, _default ...bool) bool {
	return getTyped(fm, qual, _default...)
}

// Get complex64 or default
func (fm *Navigator) Complex64(qual IQual, _default ...complex64) complex64 {
	return getTyped(fm, qual, _default...)
}

// Get complex128 or default
func (fm *Navigator) Complex128(qual IQual, _default ...complex128) complex128 {
	return getTyped(fm, qual, _default...)
}

// Get int or default
func (fm *Navigator) Int(qual IQual, _default ...int) int {
	var defaultVal int
	if len(_default) > 0 {
		defaultVal = _default[0]
	}
	if val, ok := fm.QualGet(qual); ok {
		if casted, ok := AnyToNumeric[int](val); ok {
			return casted
		}
	}
	return defaultVal
}

// Get int64 or default
func (fm *Navigator) Int64(qual IQual, _default ...int64) int64 {
	var defaultVal int64
	if len(_default) > 0 {
		defaultVal = _default[0]
	}
	if val, ok := fm.QualGet(qual); ok {
		if casted, ok := AnyToNumeric[int64](val); ok {
			return casted
		}
	}
	return defaultVal
}

// Get int32 or default
func (fm *Navigator) Int32(qual IQual, _default ...int32) int32 {
	var defaultVal int32
	if len(_default) > 0 {
		defaultVal = _default[0]
	}
	if val, ok := fm.QualGet(qual); ok {
		if casted, ok := AnyToNumeric[int32](val); ok {
			return casted
		}
	}
	return defaultVal
}

// Get uint or default
func (fm *Navigator) Uint(qual IQual, _default ...uint) uint {
	var defaultVal uint
	if len(_default) > 0 {
		defaultVal = _default[0]
	}
	if val, ok := fm.QualGet(qual); ok {
		if casted, ok := AnyToNumeric[uint](val); ok {
			return casted
		}
	}
	return defaultVal
}

// Get uint64 or default
func (fm *Navigator) Uint64(qual IQual, _default ...uint64) uint64 {
	var defaultVal uint64
	if len(_default) > 0 {
		defaultVal = _default[0]
	}
	if val, ok := fm.QualGet(qual); ok {
		if casted, ok := AnyToNumeric[uint64](val); ok {
			return casted
		}
	}
	return defaultVal
}

// Get uint32 or default
func (fm *Navigator) Uint32(qual IQual, _default ...uint32) uint32 {
	var defaultVal uint32
	if len(_default) > 0 {
		defaultVal = _default[0]
	}
	if val, ok := fm.QualGet(qual); ok {
		if casted, ok := AnyToNumeric[uint32](val); ok {
			return casted
		}
	}
	return defaultVal
}

// Get uint16 or default
func (fm *Navigator) Uint16(qual IQual, _default ...uint16) uint16 {
	var defaultVal uint16
	if len(_default) > 0 {
		defaultVal = _default[0]
	}
	if val, ok := fm.QualGet(qual); ok {
		if casted, ok := AnyToNumeric[uint16](val); ok {
			return casted
		}
	}
	return defaultVal
}

// Get uint8 or default
func (fm *Navigator) Uint8(qual IQual, _default ...uint8) uint8 {
	var defaultVal uint8
	if len(_default) > 0 {
		defaultVal = _default[0]
	}
	if val, ok := fm.QualGet(qual); ok {
		if casted, ok := AnyToNumeric[uint8](val); ok {
			return casted
		}
	}
	return defaultVal
}

// Get int16 or default
func (fm *Navigator) Int16(qual IQual, _default ...int16) int16 {
	var defaultVal int16
	if len(_default) > 0 {
		defaultVal = _default[0]
	}
	if val, ok := fm.QualGet(qual); ok {
		if casted, ok := AnyToNumeric[int16](val); ok {
			return casted
		}
	}
	return defaultVal
}

// Get int8 or default
func (fm *Navigator) Int8(qual IQual, _default ...int8) int8 {
	var defaultVal int8
	if len(_default) > 0 {
		defaultVal = _default[0]
	}
	if val, ok := fm.QualGet(qual); ok {
		if casted, ok := AnyToNumeric[int8](val); ok {
			return casted
		}
	}
	return defaultVal
}

// Get float64 or default
func (fm *Navigator) Float64(qual IQual, _default ...float64) float64 {
	var defaultVal float64
	if len(_default) > 0 {
		defaultVal = _default[0]
	}
	if val, ok := fm.QualGet(qual); ok {
		if casted, ok := AnyToNumeric[float64](val); ok {
			return casted
		}
	}
	return defaultVal
}

// Get float64 or default
func (fm *Navigator) Float32(qual IQual, _default ...float32) float32 {
	var defaultVal float32
	if len(_default) > 0 {
		defaultVal = _default[0]
	}
	if val, ok := fm.QualGet(qual); ok {
		if casted, ok := AnyToNumeric[float32](val); ok {
			return casted
		}
	}
	return defaultVal
}

func (fm *Navigator) Navigator(qual IQual) *Navigator {
	if val, ok := fm.QualGet(qual); ok {
		switch casted := val.(type) {
		case map[string]any:
			return FromMap(casted)
		case []any:
			return FromList(casted)
		case ISource:
			return New(casted)
		}
	}
	return nil
}
