package list

// ----------
// Interfaces
// ----------

// Comparable defines how two values are compared.
type Comparable interface {
	Compare(Comparable) int
}

// ---------
// Functions
// ---------

// Filterer determines if a value is to be retained.
type Filterer func(x interface{}) bool

// Generator defines the ith value in a list.
type Generator func(i int) interface{}

// Lesser defines the less-than comparison on two values.
type Lesser func(x, y interface{}) bool

// Mapper defines a value from another value.
type Mapper func(x interface{}) interface{}

// Reducer defines a value given two values.
type Reducer func(x, y interface{}) interface{}

// -------------------------------------
// Default Less function implementations
// -------------------------------------

// Bytes (type Lesser) is the less-than comparison of two interface types as bytes.
func Bytes(x, y interface{}) bool { return x.(byte) < y.(byte) }

// CmpLess (type Lesser) is the less-than comparison of two comparable types.
func CmpLess(x, y interface{}) bool { return x.(Comparable).Compare(y.(Comparable)) < 0 }

// Float32s (type Lesser) is the less-than comparison of two interface types as float32s.
func Float32s(x, y interface{}) bool { return x.(float32) < y.(float32) }

// Float64s (type Lesser) is the less-than comparison of two interface types as float64s.
func Float64s(x, y interface{}) bool { return x.(float64) < y.(float64) }

// Int8s (type Lesser) is the less-than comparison of two interface types as int8s.
func Int8s(x, y interface{}) bool { return x.(int8) < y.(int8) }

// Int16s (type Lesser) is the less-than comparison of two interface types as int16s.
func Int16s(x, y interface{}) bool { return x.(int16) < y.(int16) }

// Int32s (type Lesser) is the less-than comparison of two interface types as int32s.
func Int32s(x, y interface{}) bool { return x.(int32) < y.(int32) }

// Int64s (type Lesser) is the less-than comparison of two interface types as int64s.
func Int64s(x, y interface{}) bool { return x.(int64) < y.(int64) }

// Ints (type Lesser) is the less-than comparison of two interface types as ints.
func Ints(x, y interface{}) bool { return x.(int) < y.(int) }

// Runes (type Lesser) is the less-than comparison of two interface types as runes.
func Runes(x, y interface{}) bool { return x.(rune) < y.(rune) }

// Strings (type Lesser) is the less-than comparison of two interface types as strings.
func Strings(x, y interface{}) bool { return x.(string) < y.(string) }

// UInt8s (type Lesser) is the less-than comparison of two interface types as uint8s.
func UInt8s(x, y interface{}) bool { return x.(uint8) < y.(uint8) }

// UInt16s (type Lesser) is the less-than comparison of two interface types as uint16s.
func UInt16s(x, y interface{}) bool { return x.(uint16) < y.(uint16) }

// UInt32s (type Lesser) is the less-than comparison of two interface types as uint32s.
func UInt32s(x, y interface{}) bool { return x.(uint32) < y.(uint32) }

// UInt64s (type Lesser) is the less-than comparison of two interface types as uint64s.
func UInt64s(x, y interface{}) bool { return x.(uint64) < y.(uint64) }

// UInts (type Lesser) is the less-than comparison of two interface types as uints.
func UInts(x, y interface{}) bool { return x.(uint) < y.(uint) }
