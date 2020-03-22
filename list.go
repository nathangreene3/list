package list

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// List is a doubly-linked list.
type List struct {
	head, tail *item
	length     int
	less       Less
}

// Filterer determines if a value is to be retained.
type Filterer func(value interface{}) bool

// Generator defines the ith value in a list.
type Generator func(i int) interface{}

// Less defines the less-than comparison on two values.
type Less func(x, y interface{}) bool

// Mapper defines a value from another value.
type Mapper func(value interface{}) interface{}

// Reducer defines a value given two values.
type Reducer func(x, y interface{}) interface{}

// New list of values.
func New(less Less, values ...interface{}) *List {
	ls := List{less: less}
	return ls.Append(values...)
}

// Generate a list of n values.
func Generate(n int, gen Generator, less Less) *List {
	ls := List{less: less}
	for ; 0 < n; n-- {
		ls.InsertAt(ls.length, gen(ls.length))
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

// Get the ith value from a list. Value is not removed from the list.
func (ls *List) Get(i int) interface{} {
	return ls.get(i).value
}

// get the ith item from a list.
func (ls *List) get(i int) *item {
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
		itm := ls.get(i)
		itm.prev.next = &item{value: value, prev: itm.prev, next: itm}
		itm.prev = itm.prev.next
	}

	ls.length++
	return ls
}

// Len of a list.
func (ls *List) Len() int {
	return ls.length
}

// Less returns the default less-than comparison. Assumes less is set.
func (ls *List) Less(i, j int) bool {
	return ls.less(ls.get(i).value, ls.get(j).value)
}

// ToMap a list of values.
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

// Pop removes the tail value from a list.
func (ls *List) Pop() interface{} {
	return ls.RemoveAt(ls.length - 1)
}

// Push appends a value onto a list.
func (ls *List) Push(value interface{}) {
	ls.InsertAt(ls.length, value)
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
		itm := ls.get(i)
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

// SetLess sets the less function for a list.
func (ls *List) SetLess(less Less) *List {
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
	for itm := ls.get(i); i < j && itm != nil; itm = itm.next {
		sub.InsertAt(sub.length, itm.value)
		i++
	}

	return sub
}

// Swap two items in a list.
func (ls *List) Swap(i, j int) {
	x, y := ls.get(i), ls.get(j)
	x.value, y.value = y.value, x.value
}

// Filter ...
func (ls *List) Filter(f Filterer) *List {
	newLs := New(ls.less)
	for itm := ls.head; itm != nil; itm = itm.next {
		if f(itm.value) {
			newLs.InsertAt(newLs.length, itm.value)
		}
	}

	return newLs
}

// Map ...
func (ls *List) Map(f Mapper) *List {
	newLs := New(ls.less)
	for itm := ls.head; itm != nil; itm = itm.next {
		newLs.InsertAt(newLs.length, f(itm.value))
	}

	return newLs
}

// Reduce ...
func (ls *List) Reduce(f Reducer) interface{} {
	var value interface{}
	for itm := ls.head; itm != nil; itm = itm.next {
		value = f(value, itm.value)
	}

	return value
}
