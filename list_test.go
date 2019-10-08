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
	var (
		iters    = 2
		numItems = 5
	)

	for i := 0; i < iters; i++ {
		ls := New()
		for j := 0; j < numItems; j++ {
			ls.Insert(Int(rand.Int()))
		}

		for itm := ls.head; itm != nil && itm.next != nil; itm = itm.next {
			if 0 < itm.value.Compare(itm.next.value) {
				t.Fatalf("expected %v < %v\n", itm.value, itm.next.value)
			}
		}

		if s := ls.Slice(); len(s) != ls.length {
			t.Fatalf("expected length %d, received %d\n", ls.length, len(s))
		}
	}
}
