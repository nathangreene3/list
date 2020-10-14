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
		ls := New(func(x, y interface{}) bool { xVal, _ := x.(int); yVal, _ := y.(int); return xVal < yVal })
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
		ls = New(func(x, y interface{}) bool { xVal, _ := x.(int); yVal, _ := y.(int); return xVal < yVal }, 0, 1, 0, 2, 0)
	)

	ls.RemoveAt(4)
	ls.RemoveAt(2)
	ls.RemoveAt(0)
	ls.Append(0, 4, 0, 5, 0).Remove(0).InsertAt(2, 3).InsertAt(0, 0).Append(6)
	if len(s) != ls.Len() {
		t.Fatalf("\nexpected %v\nreceived %v\n", s, ls.String())
	}

	for i := 0; i < len(s); i++ {
		index, ok := ls.Search(s[i])
		if i != index || !ok {
			t.Fatalf("\nexpected %v\nreceived %v\n", s, ls.String())
		}
	}
}

func TestSort(t *testing.T) {
	var (
		s  = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		ls = New(func(x, y interface{}) bool { xVal, _ := x.(int); yVal, _ := y.(int); return xVal < yVal }, 9, 0, 8, 1, 7, 2, 6, 3, 5, 4).Sort()
	)

	if len(s) != ls.Len() {
		t.Fatalf("\nexpected %v\nreceived %v\n", s, ls.String())
	}

	for i := 0; i < len(s); i++ {
		index, ok := ls.Search(s[i])
		if i != index || !ok {
			t.Fatalf("\nexpected %v\nreceived %v\n", s, ls.String())
		}
	}
}

func TestReduce(t *testing.T) {
	var (
		n   = 5
		exp = n * (n + 1) / 2 // 1+2+...+n
		rec = Generate(
			n,
			func(i int) interface{} { return i + 1 },
			func(x, y interface{}) bool { xVal, _ := x.(int); yVal, _ := y.(int); return xVal < yVal },
		).Reduce(func(x, y interface{}) interface{} { xVal, _ := x.(int); yVal, _ := y.(int); return xVal + yVal })
	)

	if exp != rec {
		t.Fatalf("\nexpected %d\nreceived %d\n", exp, rec)
	}
}
