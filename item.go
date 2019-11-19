package list

import "reflect"

// item ...
type item struct {
	Value      interface{}
	prev, next *item
}

func newItem(value interface{}, prev, next *item) *item {
	return &item{Value: value, prev: prev, next: next}
}

// getFrom returns the ith item from a starting item.
func (itm *item) getFrom(i int) *item {
	for ; 0 < i; itm = itm.next {
		// TODO: Should this be allowed to panic when itm.next is nil?
		i--
	}

	return itm
}

// contains returns true if an item holds a value that is the same type and
// equal to a given value.
func (itm *item) contains(value interface{}) bool {
	v := itm.Value
	return reflect.TypeOf(v) == reflect.TypeOf(value) && v == value
}
