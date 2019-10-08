package list

import (
	"math/rand"
	"testing"
)

// Int ...
type Int int

func (n Int) Compare(m Comparable) int {
	switch {
	case n < m.(Int):
		return -1
	case m.(Int) < n:
		return 1
	default:
		return 0
	}
}

func TestList(t *testing.T) {
	var ls *List
	for i := 0; i < 10; i++ {
		ls = New()
		for j := 0; j < 10; j++ {
			ls.Insert(Int(rand.Int()))
		}

		t.Fatalf("%v\n", ls.Slice())

		for itm := ls.head; itm != nil && itm.next != nil; itm = itm.next {
			if 0 < itm.value.Compare(itm.next.value) {
				t.Fatalf("expected %v < %v\n", itm.value, itm.next.value)
			}
		}
	}
}
