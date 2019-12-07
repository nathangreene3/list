package list

import (
	"fmt"
	"strings"
)

// List is a doubly-linked list.
type List struct {
	head, tail *item
	length     int
}

// New list of values.
func New(values ...interface{}) List {
	var ls List
	ls.Append(values...)
	return ls
}

// Append several values into a list.
func (ls *List) Append(values ...interface{}) {
	for _, value := range values {
		ls.InsertAt(ls.length, value)
	}
}

// Copy a list.
func (ls *List) Copy() List {
	var cpy List
	for itm := ls.head; itm != nil; itm = itm.next {
		cpy.InsertAt(cpy.length, itm.Value)
	}

	return cpy
}

// Equal compares two lists for equality.
func (ls *List) Equal(list List) bool {
	if ls.length != list.length {
		return false
	}

	itm0, itm1 := ls.head, list.head
	for itm0 != nil && itm1 != nil {
		if !itm0.equals(itm1) {
			return false
		}

		itm0 = itm0.next
		itm1 = itm1.next
	}

	return true
}

// InsertAt inserts a value into the ith index.
func (ls *List) InsertAt(i int, value interface{}) {
	if i < 0 || ls.length < i {
		panic("index out of range")
	}

	switch i {
	case ls.length:
		if ls.length == 0 {
			// i = length = 0 --> initialize head & tail
			ls.head = &item{Value: value}
			ls.tail = ls.head
		} else {
			// 0 < i = length --> append as new tail
			ls.tail.next = &item{Value: value, prev: ls.tail}
			ls.tail = ls.tail.next
		}
	case 0:
		// 0 < length --> prepend as new head
		ls.head.prev = newItem(value, nil, ls.head)
		ls.head = ls.head.prev
	default:
		// 0 < i < length --> insert as normal
		itm := ls.head.getFrom(i)
		itm.prev.next = newItem(value, itm.prev, itm)
		itm.prev = itm.prev.next
	}

	ls.length++
}

// Join several lists into one list.
func (ls *List) Join(lists ...List) {
	for _, list := range lists {
		for itm := list.head; itm != nil; itm = itm.next {
			ls.InsertAt(ls.length, itm.Value)
		}
	}
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

// Remove all instances of a given value from a list of values.
func (ls *List) Remove(value interface{}) {
	for itm := ls.head; itm != nil; itm = itm.next {
		if itm.contains(value) {
			switch itm {
			case ls.head:
				if ls.length == 1 {
					ls.head = nil
					ls.tail = nil
				} else {
					ls.head = ls.head.next
					ls.head.prev = nil
				}
			case ls.tail:
				if ls.length == 1 {
					ls.head = nil
					ls.tail = nil
				} else {
					ls.tail = ls.tail.prev
					ls.tail.next = nil
				}
			default:
				itm.prev.next = itm.next
				itm.next.prev = itm.prev
			}

			ls.length--
		}
	}
}

// RemoveAt the ith value.
func (ls *List) RemoveAt(i int) interface{} {
	if i < 0 || ls.length <= i {
		panic("index out of range")
	}

	switch i {
	case 0:
		// Remove the head
		value := ls.head.Value
		if ls.length == 1 {
			ls.head = nil
			ls.tail = nil
		} else {
			ls.head = ls.head.next
			ls.head.prev = nil
		}

		ls.length--
		return value
	case ls.length - 1:
		// Remove the tail
		value := ls.tail.Value
		if ls.length == 1 {
			ls.head = nil
			ls.tail = nil
		} else {
			ls.tail = ls.tail.prev
			ls.tail.next = nil
		}

		ls.length--
		return value
	default:
		// Remove a normal item
		itm := ls.head.getFrom(i)
		itm.prev.next = itm.next
		itm.next.prev = itm.prev
		return itm.Value
	}
}

// Reverse a subrange of a list.
func (ls *List) Reverse(i, j int) {
	if j < i {
		i, j = j, i
	}

	itmI := ls.head.getFrom(i)
	itmJ := itmI.getFrom(j - i)
	for i < j {
		itmI.Value, itmJ.Value = itmJ.Value, itmI.Value
		itmI, itmJ = itmI.next, itmJ.prev
		i++
		j--
	}
}

// RotateLeft moves the head to the tail.
func (ls *List) RotateLeft() {
	if 1 < ls.length {
		ls.head.prev = ls.tail
		ls.tail.next = ls.head

		ls.head = ls.head.next
		ls.tail = ls.tail.next

		ls.head.prev = nil
		ls.tail.next = nil
	}
}

// RotateRight moves the tail to the head.
func (ls *List) RotateRight() {
	if 1 < ls.length {
		ls.head.prev = ls.tail
		ls.tail.next = ls.head

		ls.head = ls.head.prev
		ls.tail = ls.tail.prev

		ls.head.prev = nil
		ls.tail.next = nil
	}
}

// Search returns the index a value was found at or the length of the list and
// whether or not the value was found in the list.
func (ls *List) Search(value interface{}) (int, bool) {
	var i int
	for itm := ls.head; itm != nil; itm = itm.next {
		if itm.contains(value) {
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
		s = append(s, itm.Value)
	}

	return s
}

// String represents a formatted list.
func (ls *List) String() string {
	s := make([]string, 0, ls.length)
	for itm := ls.head; itm != nil; itm = itm.next {
		s = append(s, fmt.Sprintf("%v", itm.Value))
	}

	return "[" + strings.Join(s, " ") + "]"
}

// string2 represents a formmated list. TODO: Decide if this is better.
func (ls *List) string2() string {
	var (
		b   strings.Builder
		itm = ls.head
	)

	b.WriteString(fmt.Sprintf("[%v", itm.Value))
	for ; itm != nil; itm = itm.next {
		b.WriteString(fmt.Sprintf(", %v", itm.Value))
	}

	b.WriteByte(']')
	return b.String()
}

// Swap two values.
func (ls *List) Swap(i, j int) {
	if j < i {
		i, j = j, i
	}

	itmI := ls.head.getFrom(i)
	itmJ := itmI.getFrom(j - i)
	itmI.Value, itmJ.Value = itmJ.Value, itmI.Value
}
