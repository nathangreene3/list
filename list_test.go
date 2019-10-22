package list

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var (
	seeded bool
	seed   int64
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
	// case ts.value < u.value:
	// 	return -1
	// case u.value < ts.value:
	// 	return 1
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

func TestList(t *testing.T) {
	if !seeded {
		seed = time.Now().Unix()
		rand.Seed(seed)
		seeded = true
	}

	var (
		iters    = 100
		numItems = 100
	)

	for i := 0; i < iters; i++ {
		ls := New()
		for j := 0; j < numItems; j++ {
			ls.Insert(testInt(rand.Int()))
		}

		for itm := ls.head; itm != nil && itm.next != nil; itm = itm.next {
			if 0 < itm.Value.Compare(itm.next.Value) {
				fmt.Printf("seed: %d\n", seed)
				t.Fatalf("expected %v < %v\n", itm.Value, itm.next.Value)
			}
		}

		if s := ls.Slice(); len(s) != ls.length {
			fmt.Printf("seed: %d\n", seed)
			t.Fatalf("expected length %d, received %d\n", ls.length, len(s))
		}
	}
}

func TestRemove(t *testing.T) {
	if !seeded {
		seed = time.Now().Unix()
		rand.Seed(seed)
		seeded = true
	}

	var (
		iters    = 100
		numItems = 100
	)

	for i := 0; i < iters; i++ {
		ls := New()
		values := make([]testInt, 0, numItems)
		for j := 0; j < numItems; j++ {
			values = append(values, testInt(rand.Intn(10)))
			ls.Insert(values[j])
		}

		if s := ls.Slice(); len(s) != ls.length {
			t.Fatalf("expected length %d, received %d\n", ls.length, len(s))
		} else {
			for _, v := range s {
				ls.Remove(v)
				if ls.Contains(v) {
					fmt.Printf("seed: %d\n", seed)
					t.Fatalf("expected %v to be removed\n", v)
				}
			}
		}
	}
}

func TestList2(t *testing.T) {
	if !seeded {
		seed = time.Now().Unix()
		rand.Seed(seed)
		seeded = true
	}

	lst := New(
		&testStruct{key: 2, value: "two"},
		&testStruct{key: 4, value: "four"},
		&testStruct{key: 3, value: "three"},
		&testStruct{key: 5, value: "five"},
		&testStruct{key: 1, value: "one"},
	)

	for i := 1; i <= 5; i++ {
		if v := lst.Contains(&testStruct{key: i}); !v {
			t.Fatalf("\nexpected %t\nreceived %t\n", true, v)
		}
	}
}
