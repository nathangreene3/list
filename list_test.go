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

func initSeed() {
	if !seeded {
		seed = time.Now().Unix()
		rand.Seed(seed)
		seeded = true
	}
}

// TestList ensures manipulating a list is equivalent to manipulating a slice.
func TestList(t *testing.T) {
	initSeed()

	var (
		iters    = 100
		numItems = 100
	)

	for i := 0; i < iters; i++ {
		ls := New()
		nums := make([]int, 0, numItems)
		for j := 0; j < numItems; j++ {
			x := rand.Int()
			ls.Append(x)
			nums = append(nums, x)
		}

		var (
			itm = ls.head
			j   int
		)

		for ; j < numItems && itm != nil && itm.next != nil; itm = itm.next {
			index, ok := ls.Search(nums[j])
			if j != index || !ok {
				fmt.Printf("seed: %d\n", seed)
				t.Fatalf("\nexpected (%d, %t)\nreceived (%d, %t)\n", j, true, index, ok)
			}

			j++
		}

		if s := ls.Slice(); len(s) != ls.length {
			fmt.Printf("seed: %d\n", seed)
			t.Fatalf("\nexpected length %d\nreceived %d\n", ls.length, len(s))
		}
	}
}

func TestInsertRemove(t *testing.T) {
	var (
		s  = []int{0, 1, 2, 3, 4, 5, 6}
		ls = New(1, 2)
	)

	ls.Append(0, 4, 0, 5, 0).Remove(0).InsertAt(2, 3).InsertAt(0, 0).Append(6)
	if len(s) != ls.Length() {
		t.Fatalf("\nexpected %v\nreceived %v\n", s, ls.String())
	}

	for i := 0; i < len(s); i++ {
		index, ok := ls.Search(s[i])
		if i != index || !ok {
			t.Fatalf("\nexpected %v\nreceived %v\n", s, ls.String())
		}
	}
}
