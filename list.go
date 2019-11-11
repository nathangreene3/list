package list

import (
	"reflect"
)

// List is a doubly-linked list.
type List struct {
	head, tail *item
	length     int
}

// New list of values.
func New(values ...interface{}) *List {
	var ls List
	ls.Append(values...)
	return &ls
}

// Search returns the index a value was found at or the length of the list and
// whether or not the value was found in the list.
func (ls *List) Search(value interface{}) (int, bool) {
	var (
		i int
		t = reflect.TypeOf(value)
	)

	for itm := ls.head; itm != nil; itm = itm.next {
		if reflect.TypeOf(itm.Value) == t && value == itm.Value {
			return i, true
		}

		i++
	}

	return i, false
}

// Append several values into a list.
func (ls *List) Append(values ...interface{}) {
	for _, value := range values {
		ls.InsertAt(ls.length, value)
	}
}

// InsertAt inserts a value into the ith index.
func (ls *List) InsertAt(i int, value interface{}) {
	switch {
	case i < 0 || ls.length < i:
		panic("index out of range")
	case i == ls.length:
		if ls.length == 0 {
			// i = length = 0 --> initialize head & tail
			ls.head = &item{Value: value}
			ls.tail = ls.head
		} else {
			// 0 < i = length --> append as new tail
			ls.tail.next = &item{Value: value, prev: ls.tail}
			ls.tail = ls.tail.next
		}
	case i == 0:
		// 0 < length --> prepend as new head
		ls.head.prev = &item{Value: value, next: ls.head}
		ls.head = ls.head.prev
	case i < ls.length:
		// 0 < i < length --> insert as normal
		itm := ls.head
		for ; 0 < i && itm != nil; itm = itm.next {
			i--
		}

		itm.prev.next = &item{
			Value: value,
			prev:  itm.prev,
			next:  itm,
		}

		itm.prev = itm.prev.next
	}

	ls.length++
}

// Length of a list.
func (ls *List) Length() int {
	return ls.length
}

// Map a list of values.
func (ls *List) Map() map[int]interface{} {
	var (
		m = make(map[int]interface{})
		i int
	)

	for itm := ls.head; itm != nil; itm = itm.next {
		m[i] = itm.Value
		i++
	}

	return m
}

// RemoveAt the ith value.
func (ls *List) RemoveAt(i int) interface{} {
	switch {
	case i < 0, ls.length <= i:
		panic("index out of range")
	case i == 0:
		// Remove the head
		value := ls.head.Value
		if ls.length == 1 {
			ls.head = nil
			ls.tail = nil
		} else {
			ls.head = ls.head.next
		}

		ls.length--
		return value
	case i+1 == ls.length:
		// Remove the tail
		value := ls.tail.Value
		if ls.length == 1 {
			ls.head = nil
			ls.tail = nil
		} else {
			ls.tail = ls.tail.prev
		}

		ls.length--
		return value
	default:
		// Remove a normal item
		itm := ls.head
		for ; 0 < i && itm != nil; itm = itm.next {
			i--
		}

		itm.prev.next = itm.next
		itm.next.prev = itm.prev
		return itm.Value
	}
}

// Slice a list of values.
func (ls *List) Slice() []interface{} {
	s := make([]interface{}, 0, ls.length)
	for itm := ls.head; itm != nil; itm = itm.next {
		s = append(s, itm.Value)
	}

	return s
}
