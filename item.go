package list

import "reflect"

// item holds a value and references the previous and next items.
type item struct {
	Value      interface{}
	prev, next *item
}

// newItem returns a new item.
func newItem(value interface{}, prev, next *item) *item {
	return &item{Value: value, prev: prev, next: next}
}

// contains returns true if an item holds a value that is the same type and
// equal to a given value.
func (itm *item) contains(value interface{}) bool {
	v := itm.Value
	return reflect.TypeOf(v) == reflect.TypeOf(value) && v == value
}

// copy an item.
func (itm *item) copy() *item {
	return &item{Value: itm.Value, prev: itm.prev, next: itm.next}
}

// equals compares two items.
func (itm *item) equals(item *item) bool {
	return itm.contains(item.Value)
}

// getFrom returns the ith item from a starting item.
func (itm *item) getFrom(i int) *item {
	for ; 0 < i; itm = itm.next {
		// TODO: Should this be allowed to panic when itm.next is nil?
		i--
	}

	return itm
}
