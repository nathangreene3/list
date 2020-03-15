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
func New(values ...Comparable) *SortedList {
	var sl SortedList
	return sl.Insert(values...)
}

// Contains returns true if a value is found in a sorted list.
func (sl *SortedList) Contains(value Comparable) bool {
	return sl.find(value) != nil
}

// find an item containing a value.
func (sl *SortedList) find(value Comparable) *item {
	if 0 < sl.length {
		for itm := sl.head; itm != nil; itm = itm.next {
			r := itm.value.Compare(value)
			switch {
			case r < 0: // Continue
			case 0 < r:
				return nil
			default:
				return itm
			}
		}
	}

	return nil
}

// Insert several values.
func (sl *SortedList) Insert(values ...Comparable) *SortedList {
	for i := 0; i < len(values); i++ {
		switch {
		case sl.length == 0:
			sl.head = &item{value: values[i]}
			sl.tail = sl.head
		case 0 < sl.head.value.Compare(values[i]):
			sl.head.prev = &item{
				value: values[i],
				next:  sl.head,
			}

			sl.head = sl.head.prev
		default:
			itm := sl.tail
			for ; itm != nil && 0 < itm.value.Compare(values[i]); itm = itm.prev {
			}

			if itm == sl.tail {
				sl.tail.next = &item{
					value: values[i],
					prev:  sl.tail,
				}

				sl.tail = sl.tail.next
			} else {
				itm.next.prev = &item{
					value: values[i],
					prev:  itm,
					next:  itm.next,
				}

				itm.next = itm.next.prev
			}
		}

		sl.length++
	}

	return sl
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
		m[i] = itm.value
		i++
	}

	return m
}

// Remove several values. If duplicates exist, they will all be removed.
func (sl *SortedList) Remove(values ...Comparable) *SortedList {
	for i := 0; i < len(values); i++ {
		for itm := sl.find(values[i]); itm != nil; itm = itm.next {
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

			sl.length--
			if itm.next != nil && itm.next.value.Compare(values[i]) != 0 {
				break
			}
		}
	}

	return sl
}

// RemoveAt the ith value.
func (sl *SortedList) RemoveAt(i int) Comparable {
	if i < 0 || sl.length <= i {
		panic("index out of range")
	}

	switch i {
	case 0:
		if sl.length == 1 {
			value := sl.head.value
			sl.head = nil
			sl.tail = nil
			sl.length--
			return value
		}

		value := sl.head.value
		sl.head = sl.head.next
		sl.length--
		return value
	case sl.length - 1:
		if sl.length == 1 {
			value := sl.head.value
			sl.head = nil
			sl.tail = nil
			sl.length--
			return value
		}

		value := sl.tail.value
		sl.tail = sl.tail.prev
		sl.length--
		return value
	default:
		for itm := sl.head; itm != nil && i < sl.length; itm = itm.next {
			if i == 0 {
				itm.prev.next = itm.next
				itm.next.prev = itm.prev
				sl.length--
				return itm.value
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
		s = append(s, itm.value)
	}

	return s
}

// String represents a formatted sorted list.
func (sl *SortedList) String() string {
	s := make([]string, 0, sl.length)
	for itm := sl.head; itm != nil; itm = itm.next {
		s = append(s, fmt.Sprintf("%v", itm.value))
	}

	return strings.Join(s, ",")
}
