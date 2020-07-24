package list

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// List is a doubly-linked list. A list implements the sort and heap interface.
type List struct {
	head, tail *item
	length     int
	less       Lesser
}

// New list of values. The Less function f is optional, but is required for sorting or calling Less.
func New(f Lesser, values ...interface{}) *List {
	return (&List{less: f}).Append(values...)
}

// Generate a list of n values. The Less function f is optional, but is required for sorting or calling Less.
func Generate(n int, g Generator, f Lesser) *List {
	ls := List{less: f}
	for ; 0 < n; n-- {
		ls.InsertAt(ls.length, g(ls.length))
	}

	return &ls
}

// Append several values into a list.
func (ls *List) Append(values ...interface{}) *List {
	for i := 0; i < len(values); i++ {
		ls.InsertAt(ls.length, values[i])
	}

	return ls
}

// Copy a list.
func (ls *List) Copy() *List {
	cpy := New(ls.less)
	for itm := ls.head; itm != nil; itm = itm.next {
		cpy.InsertAt(cpy.length, itm.value)
	}

	return cpy
}

// Equal returns true if two lists contain equal values.
func (ls *List) Equal(list *List) bool {
	if ls.length != list.length {
		return false
	}

	for left, right := ls.head, list.head; left != nil && right != nil; left, right = left.next, right.next {
		if left.value != right.value {
			return false
		}
	}

	return true
}

// Filter returns a new list without the filtered values given a filter function.
func (ls *List) Filter(f Filterer) *List {
	newLs := New(ls.less)
	for itm := ls.head; itm != nil; itm = itm.next {
		if f(itm.value) {
			newLs.InsertAt(newLs.length, itm.value)
		}
	}

	return newLs
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
		itm := ls.item(i)
		itm.prev.next = &item{value: value, prev: itm.prev, next: itm}
		itm.prev = itm.prev.next
	}

	ls.length++
	return ls
}

// item returns the ith item from a list.
func (ls *List) item(i int) *item {
	if i < 0 || ls.length <= i {
		panic("index out of range")
	}

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

	return itm
}

// Len of a list.
func (ls *List) Len() int {
	return ls.length
}

// Less returns the default less-than comparison on the ith and jth items. Assumes less is set.
func (ls *List) Less(i, j int) bool {
	return ls.less(ls.item(i).value, ls.item(j).value)
}

// Map a list to a new list given a mapping function.
func (ls *List) Map(f Mapper) *List {
	newLs := New(ls.less)
	for itm := ls.head; itm != nil; itm = itm.next {
		newLs.InsertAt(newLs.length, f(itm.value))
	}

	return newLs
}

// Pop removes the tail value from a list.
func (ls *List) Pop() interface{} {
	return ls.RemoveAt(ls.length - 1)
}

// Prepend inserts values at the beginning of a list.
func (ls *List) Prepend(values ...interface{}) *List {
	for i := 0; i < len(values); i++ {
		ls.InsertAt(0, values[i])
	}

	return ls
}

// Push appends a value onto a list.
func (ls *List) Push(value interface{}) {
	ls.InsertAt(ls.length, value)
}

// Reduce a list to a value given a reducing function.
func (ls *List) Reduce(f Reducer) interface{} {
	if ls.length == 0 {
		panic("list: cannot reduce empty list")
	}

	value := ls.head.value
	for itm := ls.head.next; itm != nil; itm = itm.next {
		value = f(value, itm.value)
	}

	return value
}

// Remove values from the list.
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
	var value interface{}
	switch {
	case i < 0, ls.length <= i:
		panic("index out of range")
	case i == 0:
		// Remove the head
		value = ls.head.value
		if ls.length == 1 {
			ls.head = nil
			ls.tail = nil
		} else {
			ls.head.next.prev = nil
			ls.head = ls.head.next
		}
	case i == ls.length-1:
		// Remove the tail
		value = ls.tail.value
		if ls.length == 1 {
			ls.head = nil
			ls.tail = nil
		} else {
			ls.tail.prev.next = nil
			ls.tail = ls.tail.prev
		}
	default:
		// Remove a normal item;
		itm := ls.item(i)
		value = itm.value
		itm.prev.next = itm.next
		itm.next.prev = itm.prev
	}

	ls.length--
	return value
}

// Search returns the index a value was found at or the length of the list and
// whether or not the value was found in the list.
func (ls *List) Search(value interface{}) (int, bool) {
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

// SetLess sets the less function for a list.
func (ls *List) SetLess(less Lesser) *List {
	ls.less = less
	return ls
}

// Slice a list of values.
func (ls *List) Slice() []interface{} {
	s := make([]interface{}, 0, ls.length)
	for itm := ls.head; itm != nil; itm = itm.next {
		s = append(s, itm.value)
	}

	return s
}

// Sort a list. Assumes less is set.
func (ls *List) Sort() *List {
	sort.Sort(ls)
	return ls
}

// String represents a formatted list.
func (ls *List) String() string {
	s := make([]string, 0, ls.length<<1)
	for itm := ls.head; itm != nil; itm = itm.next {
		s = append(s, fmt.Sprintf("%v", itm.value))
	}

	return "[" + strings.Join(s, " ") + "]"
}

// SubList returns a list of the values on the range [i,j) having length j-i.
func (ls *List) SubList(i, j int) *List {
	if j < i || i < 0 || ls.length < j {
		panic("index out of range")
	}

	sub := New(ls.less)
	for itm := ls.item(i); i < j && itm != nil; itm = itm.next {
		sub.InsertAt(sub.length, itm.value)
		i++
	}

	return sub
}

// Swap two items in a list.
func (ls *List) Swap(i, j int) {
	x, y := ls.item(i), ls.item(j)
	x.value, y.value = y.value, x.value
}

// ToMap returns a map indices to their values.
func (ls *List) ToMap() map[int]interface{} {
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

// Value returns the ith value from a list. Value is not removed from the list.
func (ls *List) Value(i int) interface{} {
	return ls.item(i).value
}
