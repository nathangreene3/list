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
