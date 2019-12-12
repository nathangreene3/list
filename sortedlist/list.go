package sortedlist

import (
	"fmt"
	"strings"
)

// SortedList is a doubly linked list of sorted values.
type SortedList struct {
	head, tail *item
	length     int
}

// New creates a new sorted list of values.
func New(values ...Comparable) SortedList {
	var ls SortedList
	ls.Insert(values...)
	return ls
}

// Copy a sorted list.
func (sl *SortedList) Copy() SortedList {
	var cpy SortedList
	for itm := sl.head; itm != nil; itm = itm.next {
		cpy.Insert(itm.Value)
	}

	return cpy
}

// Contains returns true if a value is found in a sorted list.
func (sl *SortedList) Contains(value Comparable) bool {
	return sl.find(value) != nil
}

// Equal returns true if two sorted lists are equal.
func (sl *SortedList) Equal(sortedList SortedList) bool {
	if sl.length != sortedList.length {
		return false
	}

	itm0, itm1 := sl.head, sortedList.head
	for itm0 != nil && itm1 != nil {
		if !itm0.equal(itm1) {
			return false
		}

		itm0 = itm0.next
		itm1 = itm1.next
	}

	return true
}

// find an item holding a value.
func (sl *SortedList) find(value Comparable) *item {
	for itm := sl.head; itm != nil; itm = itm.next {
		c := itm.compare(value)
		switch {
		case c == 0:
			return itm
		case 0 < c:
			return nil
		}
	}

	return nil
}

// Insert several values.
func (sl *SortedList) Insert(values ...Comparable) {
	for _, value := range values {
		switch {
		case sl.length == 0:
			sl.head = newItem(value, nil, nil)
			sl.tail = sl.head
		case 0 < sl.head.Value.Compare(value):
			sl.head.prev = newItem(value, nil, sl.head)
			sl.head = sl.head.prev
		default:
			itm := sl.tail
			for ; itm != nil && 0 < itm.Value.Compare(value); itm = itm.prev {
			}

			if itm == sl.tail {
				sl.tail.next = newItem(value, sl.tail, nil)
				sl.tail = sl.tail.next
			} else {
				itm.next.prev = newItem(value, itm, itm.next)
				itm.next = itm.next.prev
			}
		}

		sl.length++
	}
}

// Join several sorted lists into one sorted list.
func (sl *SortedList) Join(sortedLists ...SortedList) {
	for _, sortedList := range sortedLists {
		for itm := sortedList.head; itm != nil; itm = itm.next {
			sl.Insert(itm.Value)
		}
	}
}

// Length of the sorted list.
func (sl *SortedList) Length() int {
	return sl.length
}

// Map comparable values.
func (sl *SortedList) Map() map[int]Comparable {
	var (
		m = make(map[int]Comparable)
		i int
	)

	for itm := sl.head; itm != nil; itm = itm.next {
		m[i] = itm.Value
		i++
	}

	return m
}

// Remove several values. If duplicates exist, they will all be removed.
func (sl *SortedList) Remove(values ...Comparable) {
	for _, value := range values {
		sl.remove(value)
	}
}

// remove a value. If duplicates exist, they will all be removed.
func (sl *SortedList) remove(value Comparable) {
	for itm := sl.find(value); itm != nil && itm.Value == value; itm = itm.next {
		switch {
		case sl.length == 1:
			sl.head = nil
			sl.tail = nil
		case itm == sl.head:
			sl.head = sl.head.next
		case itm == sl.tail:
			sl.tail = sl.tail.prev
		default:
			itm.prev.next = itm.next
			itm.next.prev = itm.prev
		}
	}
}

// RemoveAt the ith value.
func (sl *SortedList) RemoveAt(i int) Comparable {
	if i < 0 || sl.length <= i {
		panic("index out of range")
	}

	switch i {
	case 0:
		if sl.length == 1 {
			value := sl.head.Value
			sl.head = nil
			sl.tail = nil
			sl.length--
			return value
		}

		value := sl.head.Value
		sl.head = sl.head.next
		sl.length--
		return value
	case sl.length - 1:
		if sl.length == 1 {
			value := sl.head.Value
			sl.head = nil
			sl.tail = nil
			sl.length--
			return value
		}

		value := sl.tail.Value
		sl.tail = sl.tail.prev
		sl.length--
		return value
	default:
		for itm := sl.head; itm != nil && i < sl.length; itm = itm.next {
			if i == 0 {
				itm.prev.next = itm.next
				itm.next.prev = itm.prev
				sl.length--
				return itm.Value
			}

			i--
		}

		return nil
	}
}

// Slice comparable values.
func (sl *SortedList) Slice() []Comparable {
	s := make([]Comparable, 0, sl.length)
	for itm := sl.head; itm != nil; itm = itm.next {
		s = append(s, itm.Value)
	}

	return s
}

// String represents a formatted sorted list.
func (sl *SortedList) String() string {
	s := make([]string, 0, sl.length)
	for itm := sl.head; itm != nil; itm = itm.next {
		s = append(s, fmt.Sprintf("%v", itm.Value))
	}

	return "[" + strings.Join(s, " ") + "]"
}

func (sl *SortedList) string2() string {
	if sl.length == 0 {
		return "[]"
	}

	var (
		b   strings.Builder
		itm = sl.head
	)

	b.WriteString(fmt.Sprintf("[%v", itm.Value))
	itm = itm.next
	for ; itm != nil; itm = itm.next {
		b.WriteString(fmt.Sprintf(" %v", itm.Value))
	}

	b.WriteByte(']')
	return b.String()
}
