package sortedlist

import (
	"fmt"
	"math/rand"
	"testing"
)

// testInt implements Interface for testing.
type testInt int

// Compare two test ints.
func (n testInt) Compare(m Comparable) int {
	switch {
	case n < m.(testInt):
		return -1
	case m.(testInt) < n:
		return 1
	default:
		return 0
	}
}

// testStruct implements Interface for testing.
type testStruct struct {
	key   int
	value string
}

// Compare two test structs.
func (ts *testStruct) Compare(x Comparable) int {
	u := x.(*testStruct)
	switch {
	case ts.key < u.key:
		return -1
	case u.key < ts.key:
		return 1
	default:
		return 0
	}
}

func (ts *testStruct) String() string {
	return fmt.Sprintf("[%d, %s]", ts.key, ts.value)
}

func TestSortedList(t *testing.T) {
	var (
		iters    = 100
		numItems = 100
	)

	for i := 0; i < iters; i++ {
		sl := New()
		for j := 0; j < numItems; j++ {
			sl.Insert(testInt(rand.Int()))
		}

		for itm := sl.head; itm != nil && itm.next != nil; itm = itm.next {
			if 0 < itm.value.Compare(itm.next.value) {
				t.Fatalf("expected %v < %v\n", itm.value, itm.next.value)
			}
		}

		if s := sl.Slice(); len(s) != sl.length {
			t.Fatalf("expected length %d, received %d\n", sl.length, len(s))
		}
	}
}

func TestRemove(t *testing.T) {
	var (
		iters    = 100
		numItems = 100
	)

	for i := 0; i < iters; i++ {
		sl := New()
		values := make([]testInt, 0, numItems)
		for j := 0; j < numItems; j++ {
			values = append(values, testInt(rand.Intn(10)))
			sl.Insert(values[j])
		}

		if s := sl.Slice(); len(s) != sl.length {
			t.Fatalf("expected length %d, received %d\n", sl.length, len(s))
		} else {
			for i := 0; i < len(s); i++ {
				sl.Remove(s[i])
				if sl.Contains(s[i]) {
					t.Fatalf("\nexpected %v to be removed\n", s[i])
				}
			}
		}
	}
}

func TestSortedList2(t *testing.T) {
	sl := New(
		&testStruct{key: 2, value: "two"},
		&testStruct{key: 4, value: "four"},
		&testStruct{key: 3, value: "three"},
		&testStruct{key: 5, value: "five"},
		&testStruct{key: 1, value: "one"},
	)

	for i := 1; i <= 5; i++ {
		if v := sl.Contains(&testStruct{key: i}); !v {
			t.Fatalf("\nexpected %t\nreceived %t\n", true, v)
		}
	}
}
