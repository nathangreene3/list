package list

import (
	"fmt"
	"reflect"
	"strings"
)

// List is a doubly-linked list.
type List struct {
	head, tail *item
	length     int
}

// New list of values.
func New(values ...interface{}) *List {
	var ls List
	return ls.Append(values...)
}

// Append several values into a list.
func (ls *List) Append(values ...interface{}) *List {
	for i := 0; i < len(values); i++ {
		ls.InsertAt(ls.length, values[i])
	}

	return ls
}

// InsertAt inserts a value into the ith index.
func (ls *List) InsertAt(i int, value interface{}) *List {
	switch {
	case i < 0, ls.length < i:
		panic("index out of range")
	case i == ls.length:
		if ls.length == 0 {
			// i = length = 0 --> initialize head & tail
			ls.head = &item{value: value}
			ls.tail = ls.head
		} else {
			// 0 < i = length --> append as new tail
			ls.tail.next = &item{value: value, prev: ls.tail}
			ls.tail = ls.tail.next
		}
	case i == 0:
		// 0 < length --> prepend as new head
		ls.head.prev = &item{value: value, next: ls.head}
		ls.head = ls.head.prev
	default:
		// 0 < i < length --> insert as normal
		var itm *item
		if i < ls.length>>1 {
			// i is closer to 0 than n
			itm = ls.head
			for ; 0 < i && itm != nil; itm = itm.next {
				i--
			}
		} else {
			// i is closer to n than 0
			itm = ls.tail
			for ; i+1 < ls.length && itm != nil; itm = itm.prev {
				i++
			}
		}

		itm.prev.next = &item{value: value, prev: itm.prev, next: itm}
		itm.prev = itm.prev.next
	}

	ls.length++
	return ls
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
		m[i] = itm.value
		i++
	}

	return m
}

// Remove ...
func (ls *List) Remove(values ...interface{}) *List {
	for i := 0; i < len(values); i++ {
		t := reflect.TypeOf(values[i])
		for itm := ls.head; itm != nil; itm = itm.next {
			if reflect.TypeOf(itm.value) == t && values[i] == itm.value {
				switch itm {
				case ls.head:
					itm.next.prev = nil
					ls.head = itm.next
				case ls.tail:
					itm.prev.next = nil
					ls.tail = itm.prev
				default:
					itm.prev.next = itm.next
					itm.next.prev = itm.prev
				}

				ls.length--
			}
		}
	}

	return ls
}

// RemoveAt the ith value.
func (ls *List) RemoveAt(i int) interface{} {
	switch {
	case i < 0, ls.length <= i:
		panic("index out of range")
	case i == 0:
		// Remove the head
		value := ls.head.value
		if ls.length == 1 {
			ls.head = nil
			ls.tail = nil
		} else {
			ls.head.next.prev = nil
			ls.head = ls.head.next
		}

		ls.length--
		return value
	case i == ls.length-1:
		// Remove the tail
		value := ls.tail.value
		if ls.length == 1 {
			ls.head = nil
			ls.tail = nil
		} else {
			ls.tail.prev.next = nil
			ls.tail = ls.tail.prev
		}

		ls.length--
		return value
	case i < ls.length>>1:
		// Remove a normal item; i is closer to 0 than n
		itm := ls.head
		for ; 0 < i && itm != nil; itm = itm.next {
			i--
		}

		itm.prev.next = itm.next
		itm.next.prev = itm.prev
		return itm.value
	default:
		// Remove a normal item; i is closer to n than 0
		itm := ls.tail
		for ; 0 < i && itm != nil; itm = itm.prev {
			i--
		}

		itm.prev.next = itm.next
		itm.next.prev = itm.prev
		return itm.value
	}
}

// Search returns the index a value was found at or the length of the list and
// whether or not the value was found in the list.
func (ls *List) Search(value interface{}) (int, bool) {
	// TODO: start searching from both ends asynchronously?
	var (
		i int
		t = reflect.TypeOf(value)
	)

	for itm := ls.head; itm != nil; itm = itm.next {
		if reflect.TypeOf(itm.value) == t && value == itm.value {
			return i, true
		}

		i++
	}

	return i, false
}

// Slice a list of values.
func (ls *List) Slice() []interface{} {
	s := make([]interface{}, 0, ls.length)
	for itm := ls.head; itm != nil; itm = itm.next {
		s = append(s, itm.value)
	}

	return s
}

// String represents a formatted list.
func (ls *List) String() string {
	s := make([]string, 0, ls.length<<1)
	for itm := ls.head; itm != nil; itm = itm.next {
		s = append(s, fmt.Sprintf("%v", itm.value))
	}

	return "[" + strings.Join(s, " ") + "]"
}
