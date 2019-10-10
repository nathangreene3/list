package list

import (
	"fmt"
	"strings"
)

// List ...
type List struct {
	head, tail *item
	length     int
}

// New creates a new list of values.
func New(values ...Comparable) *List {
	var ls List
	ls.Insert(values...)
	return &ls
}

// Contains returns true if a value is found in a list.
func (ls *List) Contains(value Comparable) bool {
	return ls.find(value) != nil
}

// find an item holding a value.
func (ls *List) find(value Comparable) *item {
	if ls.length == 0 {
		return nil
	}

	for itm := ls.head; itm != nil; itm = itm.next {
		if itm.Value.Compare(value) == 0 {
			return itm
		}
	}

	return nil
}

// RemoveAt the ith value.
func (ls *List) RemoveAt(i int) Comparable {
	if i < 0 || ls.length <= i {
		return nil // panic("index out of range") // TODO: Maybe just return nil?
	}

	switch i {
	case 0:
		return ls.Head()
	case ls.length - 1:
		return ls.Tail()
	default:
		for itm := ls.head; itm != nil && 0 <= i; itm = itm.next {
			if i == 0 {
				itm.prev.next = itm.next
				itm.next.prev = itm.prev
				ls.length--
				return itm.Value
			}

			i--
		}

		return nil
	}
}

// Head removes the head value.
func (ls *List) Head() Comparable {
	switch ls.length {
	case 0:
		return nil
	case 1:
		value := ls.head.Value
		ls.head = nil
		ls.tail = nil
		ls.length--
		return value
	default:
		value := ls.head.Value
		ls.head = ls.head.next
		ls.length--
		return value
	}
}

// Insert several values.
func (ls *List) Insert(values ...Comparable) {
	for _, value := range values {
		ls.insert(value)
	}
}

// insert a value.
func (ls *List) insert(value Comparable) {
	switch {
	case ls.length == 0:
		ls.head = &item{Value: value}
		ls.tail = ls.head
	case 0 < ls.head.Value.Compare(value):
		ls.head.prev = &item{
			Value: value,
			next:  ls.head,
		}

		ls.head = ls.head.prev
	default:
		for itm := ls.tail; itm != nil; itm = itm.prev {
			if itm.Value.Compare(value) <= 0 {
				if itm == ls.tail {
					ls.tail.next = &item{
						Value: value,
						prev:  ls.tail,
					}

					ls.tail = ls.tail.next
				} else {
					itm.next.prev = &item{
						Value: value,
						prev:  itm,
						next:  itm.next,
					}

					itm.next = itm.next.prev
				}

				break
			}
		}
	}

	ls.length++
}

// Length of the list.
func (ls *List) Length() int {
	return ls.length
}

// Map comparable values.
func (ls *List) Map() map[int]Comparable {
	m := make(map[int]Comparable)
	var i int
	for itm := ls.head; itm != nil; itm = itm.next {
		m[i] = itm.Value
		i++
	}

	return m
}

// Remove several values. If duplicates exist, they will all be removed.
func (ls *List) Remove(values ...Comparable) {
	for _, value := range values {
		ls.remove(value)
	}
}

// remove a value. If duplicates exist, they will all be removed.
func (ls *List) remove(value Comparable) {
	for itm := ls.find(value); itm != nil; itm = ls.find(value) {
		switch {
		case ls.length == 1:
			ls.head = nil
			ls.tail = nil
		case itm == ls.head:
			ls.head = ls.head.next
		case itm == ls.tail:
			ls.tail = ls.tail.prev
		default:
			itm.prev.next = itm.next
			itm.next.prev = itm.prev
		}

		ls.length--
	}
}

// Slice comparable values.
func (ls *List) Slice() []Comparable {
	s := make([]Comparable, 0, ls.length)
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

	return strings.Join(s, ",")
}

// Tail removes the tail value.
func (ls *List) Tail() Comparable {
	switch ls.length {
	case 0:
		return nil
	case 1:
		value := ls.head.Value
		ls.head = nil
		ls.tail = nil
		ls.length--
		return value
	default:
		value := ls.tail.Value
		ls.tail = ls.tail.prev
		ls.length--
		return value
	}
}
